.PHONY: format, lint, test, run

release: check build
check: format lint test test-race clean

generate:
	docker-compose run --rm goapp go generate ./...
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
build:
	docker-compose run --rm goapp go mod tidy -v
	docker-compose run --rm goapp go build -o ./bin/app cmd/ova-template-api/main.go
