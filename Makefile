.SILENT:
.PHONY: build

all: build-docker

vendor:
	GOPROXY=https://goproxy.cn go mod vendor

build:
	docker build -t image-processing-api .
test:
	docker run -it -d -p 8080:8080 -v ./data:/data image-processing-api


