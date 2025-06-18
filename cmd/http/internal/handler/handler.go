package handler

import (
	"github.com/google/wire"
	"github.com/nuea/go-template/cmd/http/internal/handler/ping"
)

type Handlers struct {
	PingHandler *ping.Handler
}

var HandlerSet = wire.NewSet(
	ping.ProvidePingHandler,

	wire.Struct(new(Handlers), "*"),
)
