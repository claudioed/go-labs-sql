package main

import (
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logger.Info("Starting Labs SQL Application")
	logger.Info("Setting Routes")
	r, err := InitRouter()
	if err != nil {
		logger.Error("Error to configure MUX Routes")
		panic("Error to configure MUX Routes")
	}
	userRouters, userRoutersError := InitUserRouter();
	if userRoutersError != nil {
		logger.Error("Error to configure Users Routes")
		panic("Error to configure Users Routes")
	}
	userRouters.CreateUserRoutes(r)
	logger.Info("Routes Configured Successfully")
	http.ListenAndServe(":8080", r)
}
