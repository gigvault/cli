.PHONY: build test lint install clean

build:
	go build -o bin/gigvault ./cmd/gigvault

test:
	go test ./... -v

lint:
	golangci-lint run ./...

install: build
	cp bin/gigvault /usr/local/bin/

clean:
	rm -rf bin/
	go clean

