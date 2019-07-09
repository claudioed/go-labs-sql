package routes

import (
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
	"go-labs-sql/api/handlers"
)

type UserRouter struct {
	UserHandler *handlers.UserHandler
}

func (userRouter *UserRouter) CreateUserRoutes(router *mux.Router) {
	router.Handle("/v1/api/users", userRouter.UserHandler.Create()).Methods("POST")
}

func NewUserRouter(userHandler *handlers.UserHandler) *UserRouter {
	logger.Info("Creating User Router Instance")
	return &UserRouter{userHandler}
}
