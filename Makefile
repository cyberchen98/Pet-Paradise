.PHONY: server
server:
	go build -o build/server 

.PHONY: clean
clean:
	go clean -i