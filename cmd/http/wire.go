// +build wireinject

package main

import (
	"anest.top/youli/cmd/http/controller"
	"anest.top/youli/cmd/http/server"
	"anest.top/youli/config"
	"context"
	"github.com/google/wire"
)

func InitApp(ctx context.Context) (*server.App, error) {
	wire.Build(config.NewConfig, controller.Provider, server.NewServer)
	return &server.App{}, nil
}
