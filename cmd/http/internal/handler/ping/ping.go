package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func ProvidePingHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRouter(g *gin.RouterGroup) {
	g.GET("/ping", h.Ping)
}

func (h *Handler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
