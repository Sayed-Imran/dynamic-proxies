package main

import (
	"github.com/sayed-imran/dynamic-proxies/handlers"
	errorHandler "github.com/sayed-imran/dynamic-proxies/handlers/error_handler"
)

func main() {
	var microservice = handlers.Microservice{
		Name:     "app-1",
		Image:    "sayedimran/fastapi-sample-app:v4",
		Replicas: 2,
		Port:     7000,
	}
	err := microservice.CreateMicroservice()
	errorHandler.ErrorHandler(err, "Error creating microservice")
	logs:= microservice.GetMicroserviceLogs()
	println(logs)
}
