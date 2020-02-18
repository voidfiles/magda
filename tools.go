// +build tools

package tools

import (
	- "google.golang.org/grpc"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)
