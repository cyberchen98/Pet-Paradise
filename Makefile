REGISTRY   = registry.cn-shenzhen.aliyuncs.com
NAMESPACE  = golem
IMAGE_NAME = pet-paradise-server
VERSION    = dev
DOCKERFILE = deployment/Dockerfile

.PHONY: build
build:
	go build -o build/server cmd/server/*.go

.PHONY: clean
clean:
	go clean -i && rm -rf build deployment/server

.PHONY: product
product: linux docker

linux:
	GOOS=linux GOARCH=amd64 go build -o deployment/server cmd/server/*.go
docker:
	docker build -f $(DOCKERFILE) -t $(REGISTRY)/$(NAMESPACE)/$(IMAGE_NAME):$(VERSION) .
	docker push $(REGISTRY)/$(NAMESPACE)/$(IMAGE_NAME):$(VERSION)