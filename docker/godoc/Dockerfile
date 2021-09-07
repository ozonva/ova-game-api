FROM golang:1.16-alpine

RUN go install golang.org/x/tools/cmd/godoc@latest

WORKDIR /app
EXPOSE 6060

CMD godoc -http=:6060