FROM golang:1.16-alpine

RUN apk add --no-cache build-base git \
    && go get -u github.com/pressly/goose/v3/cmd/goose

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
