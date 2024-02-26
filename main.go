package main

import (
	"github.com/sayed-imran/dynamic-proxies/config"
	"github.com/sayed-imran/dynamic-proxies/handlers"
)

func main() {
	clientset := config.KubeClient.Clientset
	DynamicClient := config.KubeClient.DynamicClient
	var kubeHandler = handlers.KubernetesHandler{
		Clientset: clientset,
		DynamicClient: DynamicClient,
		Namespace: config.Configuration.Namespace,
		Name:      "myapp",
		Image:     "nginx:1.14.2",
		Replicas:  3,
		Port:      80,
	}
	err := kubeHandler.CreateDeployment()
	handlers.ErrorHandler(err, "Error creating deployment")
	err = kubeHandler.CreateService()
	handlers.ErrorHandler(err, "Error creating service")
	err = kubeHandler.CreateVirtualService()
	handlers.ErrorHandler(err, "Error creating virtual service")
}
