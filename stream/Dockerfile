FROM golang:1.23.5-alpine AS builder
WORKDIR /stream
COPY . .
RUN ls -la
RUN go build -o stream-app ./stream/stream.go

FROM alpine:latest
WORKDIR /stream
COPY --from=builder /stream/stream-app .
COPY --from=builder /stream/stream/etc ./etc
EXPOSE 8080

RUN ls -la
CMD ["./stream-app", "-f", "./etc/stream-api.yaml"]