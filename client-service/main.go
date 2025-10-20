package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	pb "github.com/pyroblazer/grpc-s3-upload/upload-service/proto"
	"google.golang.org/grpc"
)

type Config struct {
	UploadServiceHost string `json:"upload_service_host"`
}

func LoadConfig() Config {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	data, _ := os.ReadFile("config/config-" + env + ".json")
	var c Config
	_ = json.Unmarshal(data, &c)
	return c
}
func main() {
	cfg := LoadConfig()
	conn, err := grpc.Dial(cfg.UploadServiceHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial fail: %v", err)
	}
	defer conn.Close()
	client := pb.NewUploadServiceClient(conn)
	req := &pb.UploadRequest{FileName: "example.txt", FileContent: []byte("example")}
	resp, err := client.UploadFile(context.Background(), req)
	if err != nil {
		log.Fatalf("upload fail: %v", err)
	}
	log.Printf("uploaded %s msg=%s", resp.Url, resp.Message)
}
