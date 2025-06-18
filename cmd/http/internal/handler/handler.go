package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/nuea/go-template/cmd/http/internal/handler/ping"
)

type Handlers struct {
	Ping *ping.Handler
}

func (h *Handlers) RegisterRoute(g *gin.Engine) {
	rg := g.Group("/api/v1")
	{
		h.Ping.RegisterRouter(rg)
	}
}

var HandlerSet = wire.NewSet(
	ping.ProvidePingHandler,

	wire.Struct(new(Handlers), "*"),
)
