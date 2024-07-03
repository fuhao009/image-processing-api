.SILENT:
.PHONY: build build-linux

all: build-linux

vendor:
	GOPROXY=https://goproxy.cn go mod vendor

build-linux:
	docker build -t ssim-app . \
    docker run --rm ssim-app


