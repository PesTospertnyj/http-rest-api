.PHOUNY: build
build:
		go build -v ./cmd/apiserver

.PHOUNY: test		
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build		