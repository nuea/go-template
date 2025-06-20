package handler

import (
	"github.com/google/wire"
	"github.com/nuea/go-template/cmd/http/internal/handler/pingpong"
)

type Handlers struct {
	PingHandler *pingpong.Handler
}

var HandlerSet = wire.NewSet(
	pingpong.ProvidePingHandler,

	wire.Struct(new(Handlers), "*"),
)
