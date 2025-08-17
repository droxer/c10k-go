.PHONY: run_fasthttp run_gin run_echo install_deps

help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  run_fasthttp - Run the fasthttp server"
	@echo "  run_gin - Run the gin server"
	@echo "  run_echo - Run the echo server"

install_deps:
	go mod tidy
	go mod download
	go mod verify
	go mod vendor	


run_fasthttp:
	go run cmd/fasthttp/main.go

run_gin:
	go run cmd/gin/main.go

run_echo:
	go run cmd/echo/main.go
