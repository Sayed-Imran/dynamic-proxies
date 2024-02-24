package main

import (

	"github.com/sayed-imran/dynamic-proxies/config"
	"github.com/sayed-imran/dynamic-proxies/handlers"
)

func main() {

	baseConfig, err := config.LoadConfig()
	handlers.ErrorHandler(err, "Error loading config")

	var kubeHandler = handlers.KubernetesHandler{
		Namespace: baseConfig.Namespace,
		Name:      "myapp",
		Image:     "nginx:1.14.2",
		Replicas:  3,
		Port:      80,
	}
	err = kubeHandler.CreateDeployment()
	handlers.ErrorHandler(err, "Error creating deployment")
}
