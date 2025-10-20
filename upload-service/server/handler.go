package server

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	pb "github.com/pyroblazer/grpc-s3-upload/upload-service/proto"
)

type UploadServer struct {
	pb.UnimplementedUploadServiceServer
	S3Client *s3.Client
	Bucket   string
}

func NewUploadServer(region, bucket string) *UploadServer {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("AWS config err: %v", err)
	}
	return &UploadServer{S3Client: s3.NewFromConfig(cfg), Bucket: bucket}
}
func (s *UploadServer) UploadFile(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	key := fmt.Sprintf("uploads/%s", req.FileName)
	_, err := s.S3Client.PutObject(ctx, &s3.PutObjectInput{Bucket: aws.String(s.Bucket), Key: aws.String(key), Body: bytes.NewReader(req.FileContent)})
	if err != nil {
		return nil, fmt.Errorf("S3 upload fail: %v", err)
	}
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.Bucket, key)
	return &pb.UploadResponse{Url: url, Message: "File uploaded successfully"}, nil
}
