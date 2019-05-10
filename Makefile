all: build test

build:
	GOPATH="$(CURDIR)" && go install T-DFS

test: build
	GOPATH="$(CURDIR)" && go test ./src/T-DFS/
