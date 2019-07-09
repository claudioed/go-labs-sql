package routes

import "github.com/google/wire"

var RoutersSet  = wire.NewSet(NewRouter,NewUserRouter)
