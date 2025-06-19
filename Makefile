HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: tidy generate ## build for use on image
	@cd cmd/http && go build -o ../../.bin/http .

tidy: ## download dependencies 
	go mod tidy

wire: ## genarate google wire
	go run github.com/google/wire/cmd/wire ./...

generate: ## go generate
	go generate ./...

lint: ## lint
	golangci-lint run

http: ## run http server
	GIN_MODE=debug ENV=local go run cmd/http/main.go

grpc: ## run grpc server
	GO111MODULE=on ENV=local go run cmd/grpc/main.go

proto: ## generate proto files
	protoc --proto_path=proto --go_out=. --go-grpc_out=. proto/go_template/*/v1/*.proto

help: ## show this help
	@${HELP_CMD}
