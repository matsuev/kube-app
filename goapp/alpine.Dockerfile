# STAGE 1

FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app

COPY ./goapp/main.go .

COPY ./go.mod .

RUN go mod tidy

RUN go build -o goapp

# STAGE 2

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/goapp .

EXPOSE 5000

CMD ["/app/goapp"]