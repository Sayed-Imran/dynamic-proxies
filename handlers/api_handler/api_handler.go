package apihandler

import (
	apischema "github.com/sayed-imran/dynamic-proxies/web/schema"
	"github.com/sayed-imran/dynamic-proxies/handlers"
)

func CreateApp(deployReq apischema.DeployConfig) error {
	var microservice = handlers.Microservice{
		Name:     deployReq.AppName,
		Image:    deployReq.Image,
		Replicas: deployReq.Replicas,
		Port:     deployReq.Port,
	}
	err := microservice.CreateMicroservice()
	return err
}

func DeleteApp(deleteReq apischema.DeleteConfig) error {
	var microservice = handlers.Microservice{
		Name: deleteReq.AppName,
	}
	err := microservice.DeleteMicroservice()
	return err
}

func GetAppLogs(appName string) string {
	var microservice = handlers.Microservice{
		Name: appName,
	}
	logs := microservice.GetMicroserviceLogs()
	return logs
}

