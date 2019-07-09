package infra

import "github.com/google/wire"

var DbSet = wire.NewSet(NewDatabaseConnection)
