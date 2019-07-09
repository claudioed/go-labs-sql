package container

import (
	"github.com/google/wire"
	"go-labs-sql/api/handlers"
	"go-labs-sql/api/routes"
	"go-labs-sql/pkg/domain/repository"
	"go-labs-sql/pkg/infra"
)

var ApplicationSet = wire.NewSet(infra.DbSet, repository.RepositoriesSet, handlers.HandlerSet, routes.RoutersSet)
