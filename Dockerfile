# syntax=docker/dockerfile:1
FROM golang:1.16 AS builder
WORKDIR /usr/local/bin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/local/bin
COPY --from=builder /usr/local/bin .
CMD ["./app"]
