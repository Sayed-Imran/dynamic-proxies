package apihandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sayed-imran/dynamic-proxies/handlers"
	apischema "github.com/sayed-imran/dynamic-proxies/web/schema"
)

func CreateMicroservice(context *gin.Context) {
	var deployConfig apischema.DeployConfig
	if err := context.ShouldBindJSON(&deployConfig); err != nil {
		return
	}
	fmt.Println("Deploying microservice")
	fmt.Println(deployConfig)
	microservice := handlers.Microservice{
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
