# C100K in Go

## Introduction

Demonstrates scalable techniques for handling 100,000 concurrent requests in Go

## Implementation


### Standard Server

This implementation uses the standard library's `net/http` package.
```
make run_server
```

### Fasthttp Server

This implementation uses the `fasthttp` package.
```
make run_fasthttp
```

### Gin Server

This implementation uses the `gin` package.
```
make run_gin
```

