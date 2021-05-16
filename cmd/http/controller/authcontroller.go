package controller

import (
	"anest.top/youli/internal/application/handler"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Handler *handler.Handler
}

func NewAuthController(handler *handler.Handler) *AuthController {
	return &AuthController{Handler: handler}
}

func (t *AuthController) Pong() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, "pong")
	}
}
