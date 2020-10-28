package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/github"
	"github.com/dghubble/sessions"
	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
)

const (
	sessionName     = "example-github-app"
	sessionSecret   = "example cookie signing secret"
	sessionUserKey  = "githubID"
	sessionUsername = "githubUsername"
)

// sessionStore encodes and decodes session data stored in signed cookies
var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

// Config configures the main ServeMux.
type Config struct {
	GithubClientID     string
	GithubClientSecret string
}

// New returns a new ServeMux with app routes.
func New(config *Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", profileHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/confirmation", confirmation)

	// 1. Register LoginHandler and CallbackHandler
	oauth2Config := &oauth2.Config{
		ClientID:     config.GithubClientID,
		ClientSecret: config.GithubClientSecret,
		RedirectURL:  "http://localhost:8080/github/callback",
		Endpoint:     githubOAuth2.Endpoint,
	}

	// state param cookies require HTTPS by default; disable for localhost development
	stateConfig := gologin.DebugOnlyCookieConfig
	mux.Handle("/github/login", github.StateHandler(stateConfig, github.LoginHandler(oauth2Config, nil)))
	mux.Handle("/github/callback", github.StateHandler(stateConfig, github.CallbackHandler(oauth2Config, issueSession(), nil)))
	return mux
}

// issueSession issues a cookie session after successful Github login
func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		githubUser, err := github.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 2. Implement a success handler to issue some form of session
		session := sessionStore.New(sessionName)
		session.Values[sessionUserKey] = *githubUser.ID
		session.Values[sessionUsername] = *githubUser.Login
		session.Save(w)
		http.Redirect(w, req, "/profile", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

// profileHandler shows a personal profile or a login button (unauthenticated).
func profileHandler(w http.ResponseWriter, req *http.Request) {
	session, err := sessionStore.Get(req, sessionName)
	if err != nil {
		// welcome with login button
		page, _ := ioutil.ReadFile("home.html")
		fmt.Fprintf(w, string(page))
		return
	}
	// authenticated profile
	formPage, _ := ioutil.ReadFile("form.html")
	fmt.Fprintf(w, string(formPage), session.Values[sessionUsername])
}

// confirmation destroys the session on POSTs and redirects to submit page.
func confirmation(w http.ResponseWriter, req *http.Request) {

	// Retrieve and set form vars

	req.ParseForm()
	nsValue := req.FormValue("firstname") + req.FormValue("lastname") + "-" + req.FormValue("project") + "-" + req.FormValue("location")

	// Creating namespace from vars
	cmd := exec.Command("kubectl", "create", "ns", nsValue)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		print("cmd.Run() failed with %s\n", err)

		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

		tmpl := template.Must(template.ParseFiles("form-error.html"))
		data := struct {
			OutError string
		}{
			OutError: errStr,
		}
		tmpl.Execute(w, data)
		//formPage, _ := ioutil.ReadFile("form-error.html")
		//fmt.Fprintf(w, string(formPage))
		http.Redirect(w, req, "/", http.StatusFound)
	} else {
		// Redirect to confirmation page
		tmpl := template.Must(template.ParseFiles("form-confirmation.html"))
		data := struct {
			Namespace string
		}{
			Namespace: nsValue,
		}
		tmpl.Execute(w, data)

		http.Redirect(w, req, "/", http.StatusFound)
	}

}

// logoutHandler destroys the session on POSTs and redirects to home.
func logoutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		sessionStore.Destroy(w, sessionName)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

// main creates and starts a Server listening.
func main() {

	const address = "localhost:8080"
	// read credentials from environment variables if available
	config := &Config{
		GithubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	}
	// allow consumer credential flags to override config fields
	clientID := flag.String("client-id", "", "Github Client ID")
	clientSecret := flag.String("client-secret", "", "Github Client Secret")
	flag.Parse()
	if *clientID != "" {
		config.GithubClientID = *clientID
	}
	if *clientSecret != "" {
		config.GithubClientSecret = *clientSecret
	}
	if config.GithubClientID == "" {
		log.Fatal("Missing Github Client ID")
	}
	if config.GithubClientSecret == "" {
		log.Fatal("Missing Github Client Secret")
	}

	log.Printf("Starting Server listening on %s\n", address)
	err := http.ListenAndServe(address, New(config))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
