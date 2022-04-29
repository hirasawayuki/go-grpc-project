package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hirasawayuki/go-grpc-project/go-grpc-api-gateway/pkg/echo/pb"
)

type EchoRequestBody struct {
	Message string
}

func Echo(ctx *gin.Context, c pb.EchoServiceClient) {
	body := EchoRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Echo(context.Background(), &pb.EchoRequest{
		Message: body.Message,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
