ARG CONFIG_PATH

FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ./bin/playlist cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/playlist .
COPY config.env .

CMD ["./playlist"]