.PHONY: build
build:
	go build -o build/server cmd/server/*.go

.PHONY: clean
clean:
	go clean -i