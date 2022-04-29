package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hirasawayuki/go-grpc-project/go-grpc-echo-svc/pkg/config"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-echo-svc/pkg/pb"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-echo-svc/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Echo Svc on", c.Port)

	s := services.Server{}
	grpcServer := grpc.NewServer()

	pb.RegisterEchoServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
