package main

import (
	"context"
	"log"
	"time"

	"kpi/config"
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
// @host wisestone-kpi.onrender.com
// @BasePath /
// ... existing swagger comments ...

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Load().DBUri))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("kpi_system")
	router := gin.Default()

	// CORS middleware configuration
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.InitRoutes(router, db)

	log.Fatal(router.Run(":8080"))
}