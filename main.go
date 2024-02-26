package main

import (
	"github.com/sayed-imran/dynamic-proxies/config"
	"github.com/sayed-imran/dynamic-proxies/handlers"
)

func main() {
	clientset := config.KubeClient()
	var kubeHandler = handlers.KubernetesHandler{
		Clientset: clientset,
		Namespace: config.Configuration.Namespace,
		Name:      "myapp",
		Image:     "nginx:1.14.2",
		Replicas:  3,
		Port:      80,
	}
	err := kubeHandler.CreateDeployment()
	handlers.ErrorHandler(err, "Error creating deployment")
}
