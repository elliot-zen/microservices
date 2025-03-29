FROM golang:1.23.7 AS builder
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOPROXY="https://mirrors.aliyun.com/goproxy/,direct" \
    go build -a -installsuffix cgo -o order \
    ./cmd/main.go
FROM scratch
COPY --from=builder /usr/src/app/order ./order
COPY --from=builder /usr/src/app/cert ./cert
CMD ["./order"]
