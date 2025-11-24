FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY ./goapp/main.go .

COPY ./goapp/go.mod .

RUN go mod tidy

RUN go build -o goapp



FROM alpine:3.22

WORKDIR /app

COPY --from=builder ./app/goapp .

EXPOSE 8080

CMD ["/app/goapp"]