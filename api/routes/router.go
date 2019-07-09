package routes

import ("github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	logger.Info("Creating General Router")
	return router
}
