package echo

import (
	"fmt"

	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/config"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/echo/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.EchoServiceClient
}

func InitServiceClient(c *config.Config) pb.EchoServiceClient {
	cc, err := grpc.Dial(c.EchoSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect: ", err)
	}

	return pb.NewEchoServiceClient(cc)
}
