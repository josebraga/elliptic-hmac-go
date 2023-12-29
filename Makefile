GIT_ROOT ?= $(shell git rev-parse --show-toplevel)

install:
	go mod download
	go mod tidy

build: install
	mkdir -p ${GIT_ROOT}/output
	go build -o ${GIT_ROOT}/output ./cmd/...
