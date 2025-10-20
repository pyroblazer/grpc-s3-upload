package server

import (
	"encoding/json"
	"log"
	"net"
	"os"

	pb "github.com/pyroblazer/grpc-s3-upload/upload-service/proto"
	"google.golang.org/grpc"
)

type Config struct {
	AWSRegion string `json:"aws_region"`
	S3Bucket  string `json:"s3_bucket"`
	GRPCPort  string `json:"grpc_port"`
}

func LoadConfig() Config {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	path := "config/config-" + env + ".json"
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("config read err: %v", err)
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		log.Fatalf("config parse err: %v", err)
	}
	return c
}
func RunGRPCServer() {
	cfg := LoadConfig()
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUploadServiceServer(s, NewUploadServer(cfg.AWSRegion, cfg.S3Bucket))
	log.Printf("gRPC UploadService on %s bucket=%s", cfg.GRPCPort, cfg.S3Bucket)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve err: %v", err)
	}
}
