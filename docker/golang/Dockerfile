FROM golang:1.16-alpine

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
