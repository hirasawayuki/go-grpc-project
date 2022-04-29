package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/auth"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/config"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/echo"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/order"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to LoadConfig", err)
	}

	r := gin.Default()
	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	echo.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
