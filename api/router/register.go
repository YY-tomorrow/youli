package router

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	AuthCtr AuthController
}

func RegisterRouter(g *gin.Engine, c *Controller) {
	authRouter(g, c.AuthCtr)
}
