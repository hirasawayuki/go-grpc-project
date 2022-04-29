package services

import (
	"context"
	"log"

	"github.com/hirasawayuki/go-grpc-project/go-grpc-echo-svc/pkg/pb"
)

type Server struct {
	pb.UnimplementedEchoServiceServer
}

func (s *Server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Println("request echo")
	return &pb.EchoResponse{
		Message: req.Message,
	}, nil
}
