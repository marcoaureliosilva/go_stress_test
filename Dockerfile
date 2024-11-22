FROM golang:1.20 AS builder

WORKDIR /app
COPY . .

RUN go build -o carga-cli

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/carga-cli .

CMD ["./carga-cli"]
