package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type SystemConfig struct {
	HTTPPort    string `envconfig:"APP_HTTP_PORT" default:"8080"`
	GRPCPort    string `envconfig:"APP_GRPC_PORT" default:"8980"`
	ServiceName string `envconfig:"SERVICE_NAME" default:"go-template"`
}

type GoTemplateGRPCConfig struct {
	GRPCTarget     string        `envconfig:"GO_TEMPLATE_GRPC_TARGET" default:"dns://go-template"`
	RequestTimeout time.Duration `envconfig:"GO_TEMPLATE_REQUEST_TIMEOUT" default:"10s"`
}

type AppConfig struct {
	System     SystemConfig
	GoTemplate GoTemplateGRPCConfig
}

func (cfg *AppConfig) load() {
	envconfig.MustProcess("", &cfg.System)
	envconfig.MustProcess("", &cfg.GoTemplate)
}

func ProvideCofig() *AppConfig {
	env, ok := os.LookupEnv("ENV")
	if ok && env != "" {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		err := godotenv.Load(fmt.Sprintf("%v/../../.env.%v", basePath, env))
		if err != nil {
			err = godotenv.Load()
			if err != nil {
				panic(err)
			}
		}
	}
	cfg := &AppConfig{}
	cfg.load()
	return cfg
}
