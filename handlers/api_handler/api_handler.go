package apihandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	microservice_handler "github.com/sayed-imran/dynamic-proxies/handlers/microservice_handler"
	apischema "github.com/sayed-imran/dynamic-proxies/web/schema"
)

func CreateMicroservice(context *gin.Context) {
	var deployConfig apischema.DeployConfig
	if err := context.ShouldBindJSON(&deployConfig); err != nil {
		return
	}
	fmt.Println("Deploying microservice")
	microservice := microservice_handler.Microservice{
		Name:     deployConfig.AppName,
		Image:    deployConfig.Image,
		Replicas: deployConfig.Replicas,
		Port:     deployConfig.Port,
	}
	err := microservice.CreateMicroservice()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Microservice created successfully"})
}

func DeleteMicroservice(context *gin.Context) {
	var deleteConfig apischema.DeleteConfig
	if err := context.ShouldBindJSON(&deleteConfig); err != nil {
		return
	}
	fmt.Println("Deleting microservice")
	microservice := microservice_handler.Microservice{
		Name: deleteConfig.AppName,
	}
	err := microservice.DeleteMicroservice()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Microservice deleted successfully"})
}
