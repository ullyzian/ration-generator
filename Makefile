.PHONY: build
.DEFAULT_GOAL := build
build:
	go build -v ./cmd/server

migrate-local:
	migrate -database postgresql://viktor@localhost:5432/ration_generator?sslmode=disable -path migrations up