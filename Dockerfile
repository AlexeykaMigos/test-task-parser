# Build
FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download || true

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o parser ./cmd/parser

FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

RUN mkdir -p /root/output

COPY --from=builder /app/parser .

RUN chmod +x parser

ENV OUTPUT_DIR=/root/output
ENV BASE_URL=https://kuper.ru
ENV TZ=Europe/Moscow

ENTRYPOINT ["./parser"]
CMD ["-output=/root/output"]