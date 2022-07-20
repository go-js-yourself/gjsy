FROM golang:1.18-alpine3.15

WORKDIR /go/src/github.com/go-js-yourself/gjsy

COPY go.mod .
RUN go mod download
