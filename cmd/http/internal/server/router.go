package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nuea/go-template/cmd/http/internal/handler"
)

func registerRouter(gin *gin.Engine, h *handler.Handlers) {
	router := *gin.Group("/api/v1")
	{
		router.GET("/ping", h.PingHandler.Ping)
	}

}
