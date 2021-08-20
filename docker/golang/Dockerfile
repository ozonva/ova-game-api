FROM golang:1.16-alpine

RUN apk add --no-cache build-base \
    && go install github.com/golang/mock/mockgen@v1.6.0 \
    && go get github.com/golang/mock/mockgen/model

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
