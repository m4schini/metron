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
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME /src/

# Run Stage
FROM chromedp/headless-shell:latest

ENV APP_NAME miner-module

# installing curl, needed for healtcheck
RUN apt update
RUN apt install dumb-init curl -y

RUN mkdir /app
COPY --from=build /$APP_NAME /app/$APP_NAME
RUN chmod +x /app/$APP_NAME

# running init, required by headless shell to avoid zombie processes
# ENTRYPOINT ["dumb-init", "--"]

RUN ls -la /
RUN ls -la /app

ENTRYPOINT /app/$APP_NAME