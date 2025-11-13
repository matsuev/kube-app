FROM golang:1.25.4-bookworm AS builder

WORKDIR /app

COPY ./goapp_book/main.go .

COPY ./goapp_book/go.mod .

RUN go mod tidy

RUN go build -o goapp_book



FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder ./app/goapp_book .

EXPOSE 8080

CMD ["/app/goapp_book"]