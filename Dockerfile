# 使用官方的 golang 镜像作为基础镜像
FROM golang:1.19

# 设置工作目录
WORKDIR /app

# 将当前目录的所有内容复制到工作目录
COPY . .

# 使用 go mod 下载依赖
RUN go mod download

# 编译 Go 应用
RUN go build -o main .

# 暴露应用运行时需要的端口
EXPOSE 8888

# 运行应用
CMD ["./main"]
