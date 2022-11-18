.PHONY: build run build-cli lint test codegen cross-compile

dependencies:
	go mod download

build:
	go build \
		-o ./out/app \
		cmd/main.go

run:
	go run cmd/main.go

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

cross-compile:
	goreleaser --snapshot --rm-dist

docker-build:
	docker build -t hexarch-go .

docker-run:
	docker run --name hexarch-go-server -d --restart=always -p 8080:8080 hexarch-go