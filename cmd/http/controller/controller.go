package controller

import (
	"anest.top/youli/api/router"
	"anest.top/youli/internal/application/handler"
	"github.com/google/wire"
)

var Provider = wire.NewSet(handler.Provider, NewAuthController, NewController)

type Controller struct {
	AuthCtr *AuthController
}

func NewController(auth *AuthController) *router.Controller {
	return &router.Controller{AuthCtr: auth}
}
