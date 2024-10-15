module github.com/yoshihiro-shu/tech-blog-admin/backend

go 1.23.1

require (
	connectrpc.com/connect v1.17.0
	github.com/go-faker/faker/v4 v4.5.0
	golang.org/x/net v0.29.0
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
)

replace connectrpc.com/connect v1.17.0 => ./connect-go
