FROM golang:latest

RUN apt-get update \
    && apt-get install -y unzip protobuf-compiler

RUN GO111MODULE=on \
    && go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest \
    && go install github.com/golang/protobuf/protoc-gen-go@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
