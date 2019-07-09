package container

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"go-labs-sql/api/routes"
)

func InitRouter() (*mux.Router, error) {
	wire.Build(ApplicationSet)
	return nil, nil
}

func InitUserRouter() (*routes.UserRouter, error) {
	wire.Build(ApplicationSet)
	return nil, nil
}