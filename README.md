# C10K in Go

## Introduction

Demonstrates scalable techniques for handling 10,000 concurrent requests in Go

## Install dependencies

```
make install_dependencies
```

## Run Standard Server

This implementation uses the standard library's `net/http` package.
```
make run_server
```

## Run Fasthttp Server

This implementation uses the `fasthttp` package.
```
make run_fasthttp
```

## Run Gin Server

This implementation uses the `gin` package.
```
make run_gin
```

## Load Testing

Test the performance of different server implementations using `wrk`.

### Prerequisites

You need to have `wrk` installed:

- **macOS**: `brew install wrk`
- **Ubuntu/Debian**: `sudo apt-get install wrk`
- **CentOS/RHEL**: `sudo yum install wrk`
- **From source**: https://github.com/wg/wrk

### Usage

```bash
# Show load test help
make load_test

# Test individual servers
./load_test.sh standard        # Test standard library server
./load_test.sh fasthttp        # Test fasthttp server
./load_test.sh gin             # Test gin server

# Test all servers sequentially
./load_test.sh all

# Test with custom parameters
./load_test.sh standard -d 60 -c 2000 -t 16
./load_test.sh fasthttp -d 10 -c 500
./load_test.sh gin -t 4
```

### Test Parameters

- `-d, --duration`: Test duration in seconds (default: 30)
- `-c, --connections`: Total connections (default: 1000)
- `-t, --threads`: Number of threads (default: 8)

### Server Ports

- **Standard**: 8080
- **FastHTTP**: 8081
- **Gin**: 8082

### Output

Test results are saved in the `results/` directory with timestamps:
- `results/wrk_standard_YYYYMMDD_HHMMSS.txt`
- `results/wrk_fasthttp_YYYYMMDD_HHMMSS.txt`
- `results/wrk_gin_YYYYMMDD_HHMMSS.txt`

