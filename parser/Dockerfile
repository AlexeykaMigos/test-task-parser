# Build
FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download || true

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o parser main.go

# Run stage
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

RUN mkdir -p /root/output

COPY --from=builder /app/parser .
RUN chmod +x parser

ENV OUTPUT_DIR=/root/output
ENV TZ=Europe/Moscow

ENTRYPOINT ["./parser"]