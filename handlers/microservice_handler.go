package handlers

import (
	"github.com/sayed-imran/dynamic-proxies/config"
	errorHandler "github.com/sayed-imran/dynamic-proxies/handlers/error_handler"
)

type Microservice struct {
	Name     string
	Image    string
	Replicas int32
	Port     int32
}

func (m *Microservice) CreateMicroservice() error {
	var kubeHandler = KubernetesHandler{
		Clientset:     config.KubeClient.Clientset,
		DynamicClient: config.KubeClient.DynamicClient,
		Namespace:     config.Configuration.Namespace,
		Name:          m.Name,
		Image:         m.Image,
		Replicas:      m.Replicas,
		Port:          m.Port,
	}
	err := kubeHandler.CreateDeployment()
	errorHandler.ErrorHandler(err, "Error creating deployment")
	err = kubeHandler.CreateService()
	errorHandler.ErrorHandler(err, "Error creating service")
	err = kubeHandler.CreateVirtualService()
	errorHandler.ErrorHandler(err, "Error creating virtual service")
	return nil

}

func (m *Microservice) DeleteMicroservice() error {
	var kubeHandler = KubernetesHandler{
		Clientset:     config.KubeClient.Clientset,
		DynamicClient: config.KubeClient.DynamicClient,
		Namespace:     config.Configuration.Namespace,
		Name:          m.Name,
	}
	err := kubeHandler.DeleteDeployment()
	errorHandler.ErrorHandler(err, "Error deleting deployment")
	err = kubeHandler.DeleteService()
	errorHandler.ErrorHandler(err, "Error deleting service")
	err = kubeHandler.DeleteVirtualService()
	errorHandler.ErrorHandler(err, "Error deleting virtual service")
	return nil
}
