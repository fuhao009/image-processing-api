# 使用官方 Golang 镜像作为构建阶段
FROM golang:1.22 as build

# 设置工作目录
WORKDIR /app

# 安装 OpenCV 依赖
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    git \
    libgtk2.0-dev \
    pkg-config \
    libavcodec-dev \
    libavformat-dev \
    libswscale-dev \
    libjpeg-dev \
    libpng-dev \
    libtiff-dev \
    libjasper-dev \
    libdc1394-22-dev

# 安装 GoCV 和 OpenCV
RUN go get -u -d gocv.io/x/gocv
WORKDIR /go/src/gocv.io/x/gocv
RUN make install

# 将当前目录内容复制到工作目录中
COPY . .

# 构建 Go 可执行文件
RUN go build -o app .

# 使用相同的 Golang 镜像作为运行阶段
FROM golang:1.22

# 安装 OpenCV 运行时依赖
RUN apt-get update && apt-get install -y \
    libgtk2.0-0 \
    libavcodec58 \
    libavformat58 \
    libswscale5 \
    libjpeg62-turbo \
    libpng16-16 \
    libtiff5 \
    libjasper1 \
    libdc1394-22

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制文件
COPY --from=build /app/app .
COPY --from=build /app/assets /root/assets

# 运行可执行文件
CMD ["./app"]