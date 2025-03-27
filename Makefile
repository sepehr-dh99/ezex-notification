all: build test

########################################
### Tools needed for development
devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest

########################################
### Building
build:
	go build -o ./build/main ./cmd/main.go

########################################
### Testing
test:
	go test ./...

########################################
### Formatting the code
fmt:
	gofumpt -l -w .

check:
	golangci-lint run --timeout=20m0s

genproto:
	protoc --go_out ./api/grpc/proto --go_opt paths=source_relative \
    --go-grpc_out ./api/grpc/proto  --go-grpc_opt paths=source_relative \
	--proto_path=./api/grpc/proto api/grpc/proto/*.proto

.PHONY: build test
