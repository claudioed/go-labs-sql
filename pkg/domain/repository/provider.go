package repository

import "github.com/google/wire"

var RepositoriesSet  = wire.NewSet(NewPostgresUserRepository)
