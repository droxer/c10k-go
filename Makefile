.PHONY: run_server run_fasthttp

install:
	go mod tidy
	go mod download
	go mod verify
	go mod vendor

help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  run_server - Run the server"
	@echo "  run_fasthttp - Run the fasthttp server"

run_server:
	go run cmd/server/main.go

run_fasthttp:
	go run cmd/fasthttp/main.go