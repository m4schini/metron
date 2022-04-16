FROM golang:1.17.5

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o app .

CMD ["/app/app"]