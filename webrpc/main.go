package main

import (
	"net/http"

	"github.com/matthiasbruns/go-server-comparison/webrpc/helloworld"
	"github.com/matthiasbruns/go-server-comparison/webrpc/server"
)

// Go and OpenAPI
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=schema.ridl -target=golang -pkg=server -server -out=server/generated.go
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=schema.ridl -target=golang -pkg=client -client -out=client/generated.go
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=schema.ridl -target=openapi -out=docs/openapi.yml

// other languages
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=schema.ridl -target=dart -client -out=other/dart/client.generated.dart
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=schema.ridl -target=typescript -client -out=other/typescript_client/client.generated.ts
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=schema.ridl -target=typescript -server -out=other/typescript_server/server.generated.ts

func main() {
	impl := &helloworld.HelloWorldV1Service{}
	srv := server.NewHelloWorldV1Server(impl)
	srv.OnError = func(r *http.Request, rpcErr *server.WebRPCError) {
		// TODO: dealt with errors
	}
	http.ListenAndServe(":8084", srv)
}
