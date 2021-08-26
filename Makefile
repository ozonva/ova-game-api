.PHONY: format, lint, test, run

release: check build
check: format lint generate-all test-all
generate-all: generate-mock generate-proto clean
test-all: test test-race clean

format:
	docker-compose run --rm goapp go fmt ./...
lint:
	docker-compose run --rm golint golangci-lint run -v
test:
	docker-compose run --rm goapp go test -v ./...
test-race:
	docker-compose run --rm goapp go test -race -v ./...
clean:
	docker-compose run --rm goapp go clean -testcache
up:
	docker-compose up -d
down:
	docker-compose down --remove-orphans
run:
	docker-compose run --rm goapp go run cmd/ova-template-api/main.go
init-build:
	docker-compose build --pull --no-cache --parallel
init-build-%:
	docker-compose build --pull --no-cache --parallel $*
build:
	docker-compose run --rm goapp go mod tidy -v
	docker-compose run --rm goapp go build -o ./bin/app cmd/ova-template-api/main.go

generate-mock:
	docker-compose run --rm gomock go generate ./...
LOCAL_API:=pkg/hero-api
generate-proto:
	docker-compose run --rm goproto mkdir -p $(LOCAL_API)
	docker-compose run --rm goproto protoc \
		--go_out=$(LOCAL_API) --go_opt=paths=import \
		--go-grpc_out=$(LOCAL_API) --go-grpc_opt=paths=import \
		--grpc-gateway_out=$(LOCAL_API) \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=import \
		--swagger_out=allow_merge=true,merge_file_name=api:swagger \
		api/hero.proto
	docker-compose run --rm goproto mv $(LOCAL_API)/github.com/ozonva/ova-game-api/pkg/ova-game-api/* $(LOCAL_API)/
	docker-compose run --rm goproto rm -rf $(LOCAL_API)/github.com
