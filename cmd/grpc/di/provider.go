package di

import (
	"github.com/google/wire"
	"github.com/nuea/go-template/cmd/grpc/internal/server"
)

var ProviderSet = wire.NewSet(
	server.ProvideGRPCServer,
)
