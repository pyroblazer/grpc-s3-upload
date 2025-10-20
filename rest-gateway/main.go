package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/pyroblazer/grpc-s3-upload/upload-service/proto"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcEndpoint = flag.String("grpc-endpoint", "upload-service:50051", "gRPC endpoint")
		httpPort     = flag.String("http-port", ":8080", "HTTP port")
	)
	flag.Parse()
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := gw.RegisterUploadServiceHandlerFromEndpoint(ctx, mux, *grpcEndpoint, opts); err != nil {
		log.Fatalf("gateway start err: %v", err)
	}
	log.Printf("REST Gateway on %s", *httpPort)
	http.ListenAndServe(*httpPort, mux)
}
