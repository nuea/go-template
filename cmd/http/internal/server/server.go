package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/nuea/go-template/cmd/http/internal/handler"
	"github.com/nuea/go-template/internal/client"
	"github.com/nuea/go-template/internal/config"
	"github.com/oklog/run"
)

type HTTPServer struct {
	cfg    *config.AppConfig
	Gin    *gin.Engine
	Srv    *http.Server
	client *client.Client
}

func (s *HTTPServer) Serve() {
	g := &run.Group{}
	g.Add(func() error {
		s.Srv = &http.Server{
			Addr:    fmt.Sprintf(":%s", s.cfg.System.HTTPPort),
			Handler: s.Gin.Handler(),
		}
		log.Println("HTTP Server - started at ip address", s.Srv.Addr)
		return s.Srv.ListenAndServe()
	}, func(error) {
		if err := s.Srv.Shutdown(context.Background()); err != nil {
			log.Println("Failed to close HTTP server")
		}
	})

	g.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		log.Println("HTTP Server - failed")
		os.Exit(1)
	}
}

func (s *HTTPServer) load(h *handler.Handlers) {
	registerRouter(s.Gin, h)
}

func ProvideHTTPServer(cfg *config.AppConfig, h *handler.Handlers, c *client.Client) *HTTPServer {
	sv := &HTTPServer{
		cfg:    cfg,
		Gin:    gin.New(),
		Srv:    &http.Server{},
		client: c,
	}

	sv.load(h)

	sv.Gin.Use(gin.Logger())
	sv.Gin.Use(gin.Recovery())

	return sv
}
