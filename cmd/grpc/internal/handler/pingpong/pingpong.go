package pingpong

import (
	"context"
	"log"

	pingpongv1 "github.com/nuea/go-template/proto/gen/go_template/ping_pong/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcService struct {
	pingpongv1.UnimplementedPingPongServiceServer
}

func ProvidePingPongGRPCService() (pingpongv1.PingPongServiceServer, error) {
	return &grpcService{}, nil
}

func (g *grpcService) StartPingPong(ctx context.Context, req *pingpongv1.StartPingPongRequest) (*pingpongv1.StartPingPongResponse, error) {
	log.Println("GRPC request server - ", "request:", req)
	if req.Message == "" {
		return nil, status.Error(codes.InvalidArgument, "message cannot be empty")
	}

	msg := "Where are you ping?"
	if req.Message == "Ping" {
		msg = "Pong"
	}

	return &pingpongv1.StartPingPongResponse{
		Message: msg,
	}, nil
}
