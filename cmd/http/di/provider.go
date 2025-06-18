package di

import (
	"github.com/google/wire"
	"github.com/nuea/go-template/cmd/http/internal/server"
)

var ProviderSet = wire.NewSet(
	server.ProvideHTTPServer,
)
