.PHONY: build
build:
	go build -o build/server cmd/server/*.go

.PHONY: clean
clean:
	go clean -i && rm -rf build deployment/server

.PHONY: product
product:
	GOOS=linux GOARCH=amd64 go build -o deployment/server cmd/server/*.go