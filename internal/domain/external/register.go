package external

import (
	"anest.top/youli/internal/data/service"
	"anest.top/youli/internal/data/store"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(store.Provider, service.Provider, NewRegister)

type Register struct {
	AuthStore   authStore
	AuthService authService
}

func NewRegister(db *gorm.DB) *Register {
	return &Register{
		AuthStore: store.NewAuthStore(db),
	}
}
