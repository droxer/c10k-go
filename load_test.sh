#!/bin/bash
# Load test script for c10k-go servers

set -e

# Defaults
DURATION=30s
CONNECTIONS=1000
THREADS=8

# Check for wrk
if ! command -v wrk >/dev/null; then
    echo "Install wrk: brew install wrk (macOS) or apt-get install wrk (Ubuntu)"
    exit 1
fi

# Parse args
while [[ $# -gt 0 ]]; do
    case $1 in
        -d) DURATION="${2}s"; shift 2 ;;
        -c) CONNECTIONS=$2; shift 2 ;;
        -t) THREADS=$2; shift 2 ;;
        -h) echo "Usage: $0 [fasthttp|gin|all] [-d duration] [-c connections] [-t threads]"; exit 0 ;;
        *) SERVER=$1; shift ;;
    esac
done

SERVER=${SERVER:-standard}

# Run wrk on server
run_test() {
    local name=$1 port=$2
    echo "=== Testing $name on :$port ==="
    
    # Start server
    case $name in
        fasthttp) go run cmd/fasthttp/main.go & ;;
        gin) go run cmd/gin/main.go & ;;
    esac
    
    PID=$!
    
    # Wait for server
    for i in {1..30}; do
        curl -s "http://localhost:$port" >/dev/null && break
        sleep 1
    done
    
    # Run test
    mkdir -p results
    wrk -t$THREADS -c$CONNECTIONS -d$DURATION "http://localhost:$port" | tee "results/wrk_${name}_$(date +%H%M%S).txt"
    
    kill $PID 2>/dev/null || true
    wait $PID 2>/dev/null || true
    echo
}

# Execute
case $SERVER in
    "fasthttp") run_test "fasthttp" 8081 ;;
    "gin") run_test "gin" 8082 ;;
    "all") 
        run_test "fasthttp" 8081  
        run_test "gin" 8082
        ;;
    *) echo "Unknown server: $SERVER"; exit 1 ;;
esac