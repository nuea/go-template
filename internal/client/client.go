package client

import (
	"github.com/google/wire"
	gotemplate "github.com/nuea/go-template/internal/client/go_template"
)

type Client struct {
	*gotemplate.GoTemplateGRPCService
}

var ClientSet = wire.NewSet(
	gotemplate.ProvideGoTemplateServiceGRPC,
	gotemplate.ProvidePingPongServiceClient,

	wire.Struct(new(gotemplate.GoTemplateGRPCService), "*"),
	wire.Struct(new(Client), "*"),
)
