# 使用基于 Debian Bullseye 的 Golang 镜像
FROM golang:1.22-bullseye

# 设置工作目录
WORKDIR /app

# 替换镜像源为清华大学源并安装依赖
RUN echo "deb https://mirrors.tuna.tsinghua.edu.cn/debian bullseye main contrib non-free" > /etc/apt/sources.list && \
    echo "deb https://mirrors.tuna.tsinghua.edu.cn/debian bullseye-updates main contrib non-free" >> /etc/apt/sources.list && \
    echo "deb https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free" >> /etc/apt/sources.list && \
    apt-get clean && apt-get update && \
    apt-get install -y \
    build-essential \
    cmake \
    git \
    libgtk2.0-0 \
    libavcodec-dev \
    libavformat-dev \
    libswscale-dev \
    libjpeg-dev \
    libpng-dev \
    libtiff-dev \
    libjpeg62-turbo-dev \
    libdc1394-22-dev \
    pkg-config \
    unzip && \
    rm -rf /var/lib/apt/lists/*

# 复制本地的 OpenCV zip 文件
COPY opencv-4.5.2.zip /opencv/opencv-4.5.2.zip

# 安装 OpenCV
RUN cd /opencv && \
    unzip opencv-4.5.2.zip && \
    mkdir -p opencv-4.5.2/build && cd opencv-4.5.2/build && \
    cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local .. && \
    make -j$(nproc) && \
    make install && \
    ldconfig

# 设置 PKG_CONFIG_PATH 环境变量
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig:$PKG_CONFIG_PATH

# 复制项目文件到工作目录
COPY . .
RUN go mod download

# 编译 Go 项目
RUN go build -o ssim ./cmd/ssim

# 设置容器启动命令
CMD ["./ssim"]
