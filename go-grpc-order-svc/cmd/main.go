package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hirasawayuki/go-grpc-project/go-grpc-order-svc/pkg/client"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-order-svc/pkg/config"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-order-svc/pkg/db"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-order-svc/pkg/pb"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to config OrderService.", err)
	}

	h := db.Init(c.DBUrl)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)
	if err != nil {
		log.Fatalln("Failed to listing", err)
	}

	fmt.Printf("Order Service Listen. port=%v\n", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
