package di

import (
	"github.com/google/wire"
	"github.com/nuea/go-template/internal/client"
	"github.com/nuea/go-template/internal/config"
)

var InternalSet = wire.NewSet(
	ConfigSet,
	client.ClientSet,
)

var ConfigSet = wire.NewSet(
	config.ProvideCofig,
)
