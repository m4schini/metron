# Test Stage
FROM golang:1.18 as test

RUN mkdir /src
COPY . /src
WORKDIR /src

RUN go mod tidy
RUN go test ./...

# Build Stage
FROM golang:1.18-alpine as build

ENV APP_NAME miner-module

RUN mkdir /src
COPY . /src
WORKDIR /src

RUN go mod tidy
RUN go build -v -o /$APP_NAME /src/

# Run Stage
FROM alpine:latest

ENV APP_NAME miner-module

COPY --from=build /$APP_NAME .

ENTRYPOINT ./$APP_NAME