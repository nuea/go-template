package handler

import (
	"github.com/google/wire"
	"github.com/nuea/go-template/cmd/grpc/internal/handler/pingpong"
	pingpongv1 "github.com/nuea/go-template/proto/gen/go_template/ping_pong/v1"
	"google.golang.org/grpc"
)

type GrpcServices struct {
	pingpongv1.PingPongServiceServer
}

func RegisterGrpcServices(sv *grpc.Server, h *GrpcServices) {
	pingpongv1.RegisterPingPongServiceServer(sv, h)
}

var HandlerSet = wire.NewSet(
	pingpong.ProvidePingPongGRPCService,

	wire.Struct(new(GrpcServices), "*"),
)
