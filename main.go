package main

import (
	"github.com/sayed-imran/dynamic-proxies/config"
	"github.com/sayed-imran/dynamic-proxies/handlers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	baseConfig, err := config.LoadConfig()
	handlers.ErrorHandler(err, "Error loading config")

	config, err := clientcmd.BuildConfigFromFlags("", baseConfig.KubeconfigPath)
	handlers.ErrorHandler(err, "Error building kubeconfig")
	clientset, err := kubernetes.NewForConfig(config)
	handlers.ErrorHandler(err, "Error building clientset")
	var kubeHandler = handlers.KubernetesHandler{
		Clientset: clientset,
		Namespace: baseConfig.Namespace,
		Name:      "myapp",
		Image:     "nginx:1.14.2",
		Replicas:  3,
		Port:      80,
	}
	err = kubeHandler.CreateDeployment()
	handlers.ErrorHandler(err, "Error creating deployment")
}
