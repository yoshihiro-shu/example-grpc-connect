module github.com/yoshihiro-shu/tech-blog-admin/backend

go 1.23.1

require (
	connectrpc.com/connect v1.17.0
	github.com/go-faker/faker/v4 v4.5.0
	golang.org/x/net v0.29.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)

replace connectrpc.com/connect v1.17.0 => ./connect-go
