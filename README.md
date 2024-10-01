# 概要

## How to Set Up Environment

Install grpc connect into local directory for debug.

```sh
git clone git@github.com:connectrpc/connect-go.git
```

## How to Set Up Server

```sh
go run ./cmd/server/main.go
```

## How to Call

```sh
$ curl \
    --header "Content-Type: application/json" \
    --data '{}' \
    http://localhost:8080/backend.v1.BackendService/SayHello
```
