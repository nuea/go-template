package gotemplate

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/nuea/go-template/internal/config"
	pingpongv1 "github.com/nuea/go-template/proto/gen/go_template/ping_pong/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GoTemplateGRPCService struct {
	PingPongServiceClient pingpongv1.PingPongServiceClient
}

type APIClient struct {
	conn *grpc.ClientConn
}

func NewDefaultGRPCClient(target string, du time.Duration, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if opts == nil {
		opts = make([]grpc.DialOption, 0)
	}

	baseOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainStreamInterceptor(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)),
		grpc.WithIdleTimeout(du),
		grpc.WithChainUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, du)
			defer cancelFunc()
			return invoker(ctxWithTimeout, method, req, reply, cc, opts...)
		}),
	}

	return grpc.NewClient(target, append(baseOpts, opts...)...)
}

func WithRequestLoggerUnaryClient() grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Println("GRPC request client - ", "request:", req, "method:", method)
		return invoker(ctx, method, req, reply, cc, opts...)
	})
}

func ProvideGoTemplateServiceGRPC(cfg *config.AppConfig) *APIClient {
	conn, err := NewDefaultGRPCClient(cfg.GoTemplate.GRPCTarget, cfg.GoTemplate.RequestTimeout, WithRequestLoggerUnaryClient())
	if err != nil {
		panic(err)
	}

	return &APIClient{
		conn: conn,
	}
}

func ProvidePingPongServiceClient(client *APIClient) pingpongv1.PingPongServiceClient {
	return pingpongv1.NewPingPongServiceClient(client.conn)
}
