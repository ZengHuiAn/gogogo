FROM alpine:latest

# 在容器的根目录下创建 app 目录
RUN mkdir /app

# 将工作目录切换到 /app 下
WORKDIR /app

# 将微服务的服务端运行文件拷贝到 /app 下
ADD consignment-service /app/consignment-service

# 运行服务端
CMD ["./consignment-service"]