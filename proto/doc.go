package web

//go:generate go get github.com/golang/protobuf/protoc-gen-go
//go:generate bash -c "protoc -I. ./web.proto --go_out=plugins=grpc:."
