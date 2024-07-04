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

COPY . .

# 解压 OpenCV 4.9.0 源码到指定目录
RUN unzip /app/opencv.zip -d /app && \
    unzip /app/opencv_contrib.zip -d /app && \
    mv /app/opencv-4.9.0 /app/opencv && \
    mv /app/opencv_contrib-4.9.0 /app/opencv_contrib \

# 安装 OpenCV
RUN cd /opencv && \
    mkdir -p build && cd build && \
    cmake -D CMAKE_BUILD_TYPE=RELEASE \
          -D CMAKE_INSTALL_PREFIX=/usr/local \
          -D OPENCV_EXTRA_MODULES_PATH=/app/opencv_contrib/modules \
          -D ENABLE_PRECOMPILED_HEADERS=OFF \
          -D BUILD_EXAMPLES=ON \
          -D OPENCV_GENERATE_PKGCONFIG=ON .. && \
    make -j$(nproc) && \
    make install && \
    ldconfig

# 设置 PKG_CONFIG_PATH 环境变量
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig:$PKG_CONFIG_PATH

# 编译 Go 项目
RUN PKG_CONFIG_PATH=/usr/local/lib/pkgconfig go build -buildvcs=false -o image-processing-api .

# 设置容器启动命令
CMD ["./ssim"]
