package pingpong

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuea/go-template/internal/client"
	pingpongv1 "github.com/nuea/go-template/proto/gen/go_template/ping_pong/v1"
)

type Handler struct {
	pingpongclient pingpongv1.PingPongServiceClient
}

func ProvidePingHandler(c *client.Client) *Handler {
	return &Handler{
		pingpongclient: c.PingPongServiceClient,
	}
}

func (h *Handler) PingPong(ctx *gin.Context) {
	var req *PingPongRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.pingpongclient.StartPingPong(ctx, &pingpongv1.StartPingPongRequest{
		Message: req.Message,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, PingPongResponse{
		Message: res.Message,
	})
}
