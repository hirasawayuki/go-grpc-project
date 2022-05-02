package db

import (
	"log"

	"github.com/hirasawayuki/go-grpc-project/go-grpc-order-svc/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Order{})
	return Handler{db}
}
