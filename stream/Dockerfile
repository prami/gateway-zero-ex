FROM golang:1.23.5-alpine AS builder
WORKDIR /stream
COPY . .
RUN ls -la
RUN go build -o stream-app stream/stream.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /stream/stream .
COPY --from=builder /stream/stream/etc /stream/etc
EXPOSE 8080
CMD ["./stream", "-f", "/stream/etc/stream-api.yaml"]