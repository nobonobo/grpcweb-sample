module github.com/nobonobo/grpcweb-sample

go 1.12

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.2.0
	github.com/gopherjs/vecty v0.0.0-20190701174234-2b6fc20f8913
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/gowebapi/webapi v0.0.0-20190324064807-2e523643cc89
	github.com/improbable-eng/grpc-web v0.9.6
	github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223 // indirect
	github.com/rs/cors v1.6.0 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.22.0
)

replace google.golang.org/grpc => github.com/johanbrandhorst/grpc-go v1.2.1-0.20180625151142-1f109e898476
