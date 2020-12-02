.PHONY: build
.DEFAULT_GOAL := build
build:
	go build -v ./cmd/server
