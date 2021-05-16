package store

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewMysql, NewAuthStore)
