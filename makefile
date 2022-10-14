.PHONY: build run build-cli lint test codegen

dependencies:
	go mod download

build:
	go build \
		-o ./out/app \
		cmd/http/main.go

run:
	go run cmd/http/main.go

build-cli:
	go build \
		-o ./out/cli \
		cmd/cli/main.go

lint:
	golangci-lint run --issues-exit-code 0 --out-format code-climate | jq -c '.[] | select(.severity|contains("major"))'

test:
	go test -v ./internal/...

codegen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1
	oapi-codegen -generate types -package dto ./docs/openapi.yaml > ./internal/core/dtos/api_dtos.gen.go
	oapi-codegen -generate spec -package http ./docs/openapi.yaml > ./internal/core/http/http_spec.gen.go