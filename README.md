# 概要

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
