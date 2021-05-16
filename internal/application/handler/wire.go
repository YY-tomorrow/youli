package handler

import (
	"anest.top/youli/internal/domain/external"
	"github.com/google/wire"
)

var Provider = wire.NewSet(external.Provider, NewHandler)
