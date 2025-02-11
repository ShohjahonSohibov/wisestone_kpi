package main

import (
	"context"
	"log"
	"time"

	_ "kpi/docs"
	"kpi/internal/app"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title KPI System API
// @version 1.0
// @description API for KPI system for Wisestone T Company
// @contact.name Developer Team
// @contact.email support@wisestonet.com
// @host localhost:8080
// @BasePath /

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("kpi_system")
	router := gin.Default()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.InitRoutes(router, db)

	log.Fatal(router.Run(":8080"))
}
