package handlers

import "fmt"

func ErrorHandler(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
	}
}
