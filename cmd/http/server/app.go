package server

import (
	"anest.top/youli/api/proto/config"
	"anest.top/youli/api/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type App struct {
	e *gin.Engine
}

func NewServer(ctx context.Context, c *router.Controller, config *config.Config) (e *App, err error) {
	serverConf := config.ServerConf
	engine := gin.Default()
	router.RegisterRouter(engine, c)
	server := &http.Server{
		Addr:           serverConf.Host,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		<-ctx.Done()
		server.Shutdown(ctx)
	}()

	return &App{engine}, server.ListenAndServe()
}
