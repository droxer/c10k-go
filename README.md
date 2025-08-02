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

