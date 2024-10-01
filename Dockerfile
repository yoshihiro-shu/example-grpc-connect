# ベースイメージとして公式の Golang イメージを使用
FROM golang:1.23.1

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY ./cmd ./cmd
COPY ./protobuf ./protobuf
COPY ./service ./service
COPY ./interceptor ./interceptor

EXPOSE 2345 8080

# アプリケーションをビルド
RUN ["go", "build", "-o", "tmp/main", "./cmd/server/main.go"]

# コンテナ起動時に実行するコマンド
# CMD ["./main"]
# CMD ["go", "run", "./cmd/server/main.go"]

# Install air & delve
RUN go install github.com/cosmtrek/air@v1.42.0 && \
    go install github.com/go-delve/delve/cmd/dlv@latest

COPY .air.toml ./

CMD ["air", "-c", ".air.toml" ]
