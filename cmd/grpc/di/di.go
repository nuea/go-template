//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nuea/go-template/cmd/grpc/internal/handler"
	internalDI "github.com/nuea/go-template/internal/di"
)

var MainSet = wire.NewSet(
	internalDI.InternalSet,
	ProviderSet,
	handler.HandlerSet,

	wire.Struct(new(Container), "*"),
)

func InitContainer() (*Container, func(), error) {
	wire.Build(MainSet)

	return &Container{}, func() {}, nil
}
