default: clean test build

test:
	go test -v ./...
.PHONY: build

build:
	go build -o server
.PHONY: build

clean:
	rm server
.PHONY: clean

fmt:
	go fmt ./...
.PHONY: fmt

run:
	go run .
