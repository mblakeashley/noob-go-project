package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	//Checking Cluster Status
	out, err := exec.Command("kubectl", "cluster-info", "--context", "kind-kind").Output()
	if err != nil {
		fmt.Println("Kind cluster does not exist, will deploy cluster...")
		output := string(out[:])
		fmt.Println(output)

		// Deploy Kind Cluster with predefined NodePort values
		cmd := exec.Command("kind", "create", "cluster", "--config=cluster-ingress.yml")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		} else {
			fmt.Println("\n" + "+ -------- Starting Infra Setup -------- +")
			fmt.Println("- Running Make Build...")
			out, err := exec.Command("make", "build").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Make Build Complete!")
			}
			fmt.Println("- Running Terraform Init... ")
			out, err = exec.Command("make", "init").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Init Complete!")
			}
			fmt.Println("- Running Terraform Plan... ")
			out, err = exec.Command("make", "plan").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Plan Complete!")
			}
			fmt.Println("- Running Terraform Apply... (Grab a cup of coffee, this will take a few minutes to complete.) ")
			out, err = exec.Command("make", "apply").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Apply Complete!")
			}
			fmt.Println("- Applying Kube Patch Configs... ")
			out, err = exec.Command("sh", "infra/scripts/kubectl_patch.sh").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Kube Patch Complete!")
				fmt.Println("+ -------- Infra Setup Complete -------- +" + "\n")
				fmt.Println("Run: 'kubectl get pods --all-namespaces' to view pods that are running in the cluster")
			}
		}

	} else {
		fmt.Println("Kind cluster exists, will delete cluster...")
		out, err = exec.Command("kind", "delete", "cluster").Output()
		if err != nil {
			fmt.Printf("%s", err)
		}

		fmt.Println("Deploying New Cluster... \n ")

		// Deploy Kind Cluster with predefined NodePort values
		cmd := exec.Command("kind", "create", "cluster", "--config=cluster-ingress.yml")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		} else {
			fmt.Println("\n" + "+ -------- Starting Infra Setup -------- +")
			fmt.Println("- Running Make Build...")
			out, err := exec.Command("make", "build").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Make Build Complete!")
			}
			fmt.Println("- Running Terraform Init... ")
			out, err = exec.Command("make", "init").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Init Complete!")
			}
			fmt.Println("- Running Terraform Plan... ")
			out, err = exec.Command("make", "plan").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Plan Complete!")
			}
			fmt.Println("- Running Terraform Apply... (Grab a cup of coffee, this will take a few minutes to complete.) ")
			out, err = exec.Command("make", "apply").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Apply Complete!")
			}
			fmt.Println("- Applying Kube Patch Configs... ")
			out, err = exec.Command("sh", "infra/scripts/kubectl_patch.sh").Output()
			if err != nil {
				output := string(out[:])
				fmt.Println(output)
			} else {
				fmt.Println("+ Kube Patch Complete!")
				fmt.Println("+ -------- Infra Setup Complete -------- +" + "\n")
				fmt.Println("Run: 'kubectl get pods --all-namespaces' to view pods that are running in the cluster")
			}
		}

	}

}
