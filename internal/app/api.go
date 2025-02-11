package app

import (
	"kpi/internal/handlers"
	"kpi/internal/repositories"
	"kpi/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(router *gin.Engine, db *mongo.Database) {
	// Initialize managers
	repoManager := repositories.NewRepositoryManager(db)
	serviceManager := services.NewServiceManager(repoManager)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(serviceManager.AuthService)
	userHandler := handlers.NewUserHandler(serviceManager.UserService)

	// API routes
	api := router.Group("/api/v1") // Add version for better API management
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			// Add other auth routes like register, logout etc.
		}

		// User routes
		users := api.Group("/users")
		{
			users.GET("", userHandler.ListUsers)
			users.GET("/:email", userHandler.GetUser)
			users.POST("", userHandler.CreateUser)
			users.PUT(":id", userHandler.UpdateUser)
			// Add DELETE endpoint for completeness
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}