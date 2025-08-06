.PHONY: run_fasthttp run_gin install_dependencies load_test

help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  run_fasthttp - Run the fasthttp server"
	@echo "  run_gin - Run the gin server"
	@echo "  load_test - Run load tests"

install_dependencies:
	go mod tidy
	go mod download
	go mod verify
	go mod vendor	


run_fasthttp:
	go run cmd/fasthttp/main.go

run_gin:
	go run cmd/gin/main.go

load_test:
	@echo "Running load tests..."
	@./load_test.sh --help