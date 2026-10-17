# 构建阶段：编译静态二进制
FROM golang:1.26-alpine AS builder

WORKDIR /src

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o /out/gowebunit .

# 运行阶段：只保留二进制
FROM alpine:3.23

RUN adduser -D -u 65532 app

WORKDIR /app
COPY --from=builder /out/gowebunit /app/gowebunit

USER app

EXPOSE 9909

ENTRYPOINT ["/app/gowebunit"]
CMD ["-port", "9909"]
