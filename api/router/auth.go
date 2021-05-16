package router

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Pong() gin.HandlerFunc
}

func authRouter(g *gin.Engine, c AuthController) {
	g.GET("/ping", c.Pong())
}
