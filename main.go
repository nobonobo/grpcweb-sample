package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/nobonobo/grpcweb-sample/backend"
	web "github.com/nobonobo/grpcweb-sample/proto"
)

const TLS = true

func main() {
	gs := grpc.NewServer()
	web.RegisterBackendServer(gs, &backend.Backend{})
	wrappedServer := grpcweb.WrapServer(gs)

	rev := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.ProtoMajor, r.Method, r.URL.Path)
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			log.Println("connect grpc:", r.RemoteAddr)
			wrappedServer.ServeHTTP(w, r)
			return
		}
		rev.ServeHTTP(w, r)
	})
	log.Println("http serve")
	if TLS {
		if err := http.ListenAndServeTLS("", "cert.pem", "key.pem", nil); err != nil {
			log.Fatal(err)
		}
	}
	if err := http.ListenAndServe("", nil); err != nil {
		log.Fatal(err)
	}
}
