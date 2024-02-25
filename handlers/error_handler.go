package handlers

import "fmt"

func ErrorHandler(err error, message string) {
	if err != nil {
		panic(fmt.Errorf("%s: %v", message, err))
	}
}
