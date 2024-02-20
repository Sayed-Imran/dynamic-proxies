package main

import (
	"fmt"

	"github.com/sayed-imran/dynamic-proxies/handlers"
)

func main() {
	var kubeHandler = handlers.KubernetesHandler{
		Namespace: "default",
		Name:      "myapp",
		Image:     "nginx:1.14.2",
		Replicas:  3,
		Port:      80,
	}
	err := kubeHandler.CreateDeployment()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
