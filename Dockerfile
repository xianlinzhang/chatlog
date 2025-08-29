# 构建阶段
FROM golang:1.24.3-alpine3.20 AS builder

# 配置环境变量
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GO111MODULE=on

# 安装依赖工具 (air编译需gcc)
RUN apk add --no-cache build-base git

# 安装热重载工具
RUN go install github.com/cosmtrek/air@v1.51.0

# 最终阶段
FROM golang:1.24.3-alpine3.20

# 从构建阶段复制air
COPY --from=builder /go/bin/air /usr/local/bin/air

# 创建工作目录
WORKDIR /app

# 暴露调试端口
EXPOSE 40000 8080

# 在最终阶段添加PATH检查（调试用）
RUN echo "Final PATH: $PATH" && \
    echo "Air location:" && \
    which air


# 启动命令 (默认运行air)
#ENTRYPOINT ["air"]