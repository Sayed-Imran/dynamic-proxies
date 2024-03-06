package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sayed-imran/dynamic-proxies/config"
	apihandler "github.com/sayed-imran/dynamic-proxies/handlers/api_handler"
)

func main() {
	// var microservice = handlers.Microservice{
	// 	Name:     "app-1",
	// 	Image:    "sayedimran/fastapi-sample-app:v4",
	// 	Replicas: 2,
	// 	Port:     7000,
	// }
	// err := microservice.CreateMicroservice()
	// errorHandler.ErrorHandler(err, "Error creating microservice")
	// logs:= microservice.GetMicroserviceLogs()
	// println(logs)
	router := gin.Default()
	router.POST("/create-microservice", apihandler.CreateMicroservice)
	router.DELETE("/delete-microservice", apihandler.DeleteMicroservice)
	router.Run(fmt.Sprintf(":%s", config.Configuration.Port))
}
