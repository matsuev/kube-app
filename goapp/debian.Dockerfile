# STAGE 1

FROM golang:1.25.4-bookworm AS builder

WORKDIR /app

COPY ./goapp/main.go .

COPY ./go.mod .

RUN go mod tidy

RUN go build -o goapp

# STAGE 2

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/goapp .

EXPOSE 5000

CMD ["/app/goapp"]