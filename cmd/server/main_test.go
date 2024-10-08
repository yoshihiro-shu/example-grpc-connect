package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"connectrpc.com/connect"
	"github.com/go-faker/faker/v4"
	"github.com/yoshihiro-shu/tech-blog-admin/backend/interceptor"
	backendv1 "github.com/yoshihiro-shu/tech-blog-admin/backend/protobuf/gen/backend/v1"        // generated by protoc-gen-go
	"github.com/yoshihiro-shu/tech-blog-admin/backend/protobuf/gen/backend/v1/backendv1connect" // generated by protoc-gen-connect-go
	"github.com/yoshihiro-shu/tech-blog-admin/backend/service"
)

var endpoint string
var httpServer *httptest.Server

func TestMain(m *testing.M) {
	srv := &service.BackendServer{}
	interceptors := connect.WithInterceptors(interceptor.Logger())

	mux := http.NewServeMux()
	path, handler := backendv1connect.NewBackendServiceHandler(srv, interceptors)
	mux.Handle(path, handler)

	httpServer = httptest.NewTLSServer(mux)
	httpServer.EnableHTTP2 = true
	defer httpServer.Close()

	endpoint = httpServer.URL

	log.Printf("server is started at %s", endpoint)

	os.Exit(m.Run())
}

func TestHttpRequest(t *testing.T) {
	// リクエストボディの作成
	requestBody, err := json.Marshal(backendv1.SayHelloRequest{
		Name: faker.Name(),
	})

	if err != nil {
		t.Log("JSONエンコードエラー:", err)
		return
	}

	// HTTPリクエストの作成
	url := endpoint + "/backend.v1.BackendService/SayHello"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Log("リクエスト作成エラー:", err)
		return
	}

	// ヘッダーの設定
	req.Header.Set("Content-Type", "application/json")

	// HTTPクライアントの作成と実行
	resp, err := httpServer.Client().Do(req)
	if err != nil {
		t.Log("リクエスト送信エラー:", err)
		return
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Log("レスポンス読み取りエラー:", err)
		return
	}
	// レスポンスの表示
	t.Logf(resp.Proto)
	t.Log(string(body))
}

func TestGRPCWebClientRequest(t *testing.T) {
	ctx := context.Background()
	opts := []connect.ClientOption{
		connect.WithGRPCWeb(),
	}
	client := backendv1connect.NewBackendServiceClient(httpServer.Client(), endpoint, opts...)
	res, err := client.SayHello(ctx, &connect.Request[backendv1.SayHelloRequest]{
		Msg: &backendv1.SayHelloRequest{
			Name: faker.Name(),
		},
	})
	if err != nil {
		t.Fatalf("error %s", err)
	}
	t.Logf("res: %#v", res)
}

func TestGRPCClientRequest(t *testing.T) {
	ctx := context.Background()
	opts := []connect.ClientOption{
		connect.WithGRPC(),
	}
	client := backendv1connect.NewBackendServiceClient(httpServer.Client(), endpoint, opts...)
	res, err := client.SayHello(ctx, &connect.Request[backendv1.SayHelloRequest]{
		Msg: &backendv1.SayHelloRequest{
			Name: faker.Name(),
		},
	})
	if err != nil {
		t.Fatalf("error %s", err)
	}
	t.Logf("res: %#v", res)
}
