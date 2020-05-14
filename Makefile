REGISTRY   = registry.cn-shenzhen.aliyuncs.com
NAMESPACE  = golem
IMAGE_NAME = pet-paradise-server
VERSION    = dev
DOCKERFILE = deployment/Dockerfile

.PHONY: build
build: server agent

server:
	go build -o build/server cmd/server/*.go

agent:
	go build -o build/agent cmd/agent/*.go

.PHONY: product
product: linux docker

linux:
	GOOS=linux GOARCH=amd64 go build -o deployment/server cmd/server/*.go
	GOOS=linux GOARCH=amd64 go build -o deployment/agent cmd/agent/*.go
docker:
	docker build -f $(DOCKERFILE) -t $(REGISTRY)/$(NAMESPACE)/$(IMAGE_NAME):$(VERSION) .
	docker push $(REGISTRY)/$(NAMESPACE)/$(IMAGE_NAME):$(VERSION)

.PHONY: clean
clean:
	go clean -i && rm -rf build deployment/server deployment/agent
