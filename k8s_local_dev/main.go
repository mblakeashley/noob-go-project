package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	//Checking Cluster Status
	out, err := exec.Command("kubectl", "cluster-info", "--context", "kind-kind").Output()
	if err != nil {
		fmt.Println("Kind cluster does not exist, will deploy cluster ")
		output := string(out[:])
		fmt.Println(output)

		createCluster := exec.Command("kind", "create", "cluster")

		createCluster.Stdout = os.Stdout
		createCluster.Stderr = os.Stdout

		if err := createCluster.Run(); err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Kind cluster exists, will delete cluster ")
		out, err := exec.Command("kind", "delete", "cluster").Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Println("Deploying new cluster.. ")
		output := string(out[:])
		fmt.Println(output)

		createCluster := exec.Command("kind", "create", "cluster")

		createCluster.Stdout = os.Stdout
		createCluster.Stderr = os.Stdout

		if err := createCluster.Run(); err != nil {
			fmt.Println("Error:", err)
		}

	}

}
