# Test Stage
FROM golang:1.18 as test

RUN mkdir /src
COPY . /src
WORKDIR /src

RUN go mod tidy
RUN go test ./...

# Build Stage
FROM golang:1.18-alpine as build

ENV APP_NAME commander-module

RUN mkdir /src
COPY . /src
WORKDIR /src

RUN go mod tidy
RUN go build -v -o /$APP_NAME /src/

# Run Stage
FROM alpine:latest

ENV APP_NAME commander-module

COPY --from=build /$APP_NAME .

RUN mkdir /config

COPY ./config/metron.yaml /config/metron.yaml

ENTRYPOINT ./$APP_NAME