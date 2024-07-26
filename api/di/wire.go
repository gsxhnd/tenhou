//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/tenhou/api/handler"
	"github.com/gsxhnd/tenhou/api/middleware"
	"github.com/gsxhnd/tenhou/api/router"
	"github.com/gsxhnd/tenhou/api/service"
)

func InitApp() (*Application, error) {
	wire.Build(
		// utils.UtilsSet,
		NewApplication,
		router.NewRouter,
		middleware.NewMiddleware,
		handler.HandlerSet,
		service.ServiceSet,
		// db.DBSet,
	)
	return &Application{}, nil
}
