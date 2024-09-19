# 基础镜像
FROM golang:1.22 as builder

# 设置工作目录
WORKDIR /app

# 复制项目代码到容器中
COPY . .

# 编译二进制文件
RUN go build -o myapp .

# 使用更小的基础镜像来运行编译后的文件
FROM alpine:3.18

# 设置工作目录
WORKDIR /app

# 从编译阶段复制二进制文件
COPY --from=builder /app/main .

# 启动应用
CMD ["./main"]
