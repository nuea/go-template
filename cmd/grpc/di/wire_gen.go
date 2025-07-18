// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nuea/go-template/cmd/grpc/internal/handler"
	"github.com/nuea/go-template/cmd/grpc/internal/handler/pingpong"
	"github.com/nuea/go-template/cmd/grpc/internal/server"
	"github.com/nuea/go-template/internal/config"
	"github.com/nuea/go-template/internal/di"
)

// Injectors from di.go:

func InitContainer() (*Container, func(), error) {
	appConfig := config.ProvideCofig()
	pingPongServiceServer, err := pingpong.ProvidePingPongGRPCService()
	if err != nil {
		return nil, nil, err
	}
	grpcServices := &handler.GrpcServices{
		PingPongServiceServer: pingPongServiceServer,
	}
	grpcServer := server.ProvideGRPCServer(appConfig, grpcServices)
	container := &Container{
		server: grpcServer,
	}
	return container, func() {
	}, nil
}

// di.go:

var MainSet = wire.NewSet(di.InternalSet, ProviderSet, handler.HandlerSet, wire.Struct(new(Container), "*"))
