package echo

import (
	"github.com/gin-gonic/gin"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/config"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/echo/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/echo")
	routes.POST("/", svc.Echo)
}

func (svc *ServiceClient) Echo(ctx *gin.Context) {
	routes.Echo(ctx, svc.Client)
}
