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
	teamHandler := handlers.NewTeamHandler(serviceManager.TeamService)

	// Initialize role handler
	roleHandler := handlers.NewRoleHandler(serviceManager.RoleService)

	// Initialize permission handler
	permissionHandler := handlers.NewPermissionHandler(serviceManager.PermissionService)

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

		teams := api.Group("/teams")
		{
			teams.GET("", teamHandler.ListTeams)
			teams.GET("/:id", teamHandler.GetTeam)
			teams.POST("", teamHandler.CreateTeam)
			teams.PUT("/:id", teamHandler.UpdateTeam)
			teams.DELETE("/:id", teamHandler.DeleteTeam)
		}

		// Role routes
		roles := api.Group("/roles")
		{
			roles.GET("", roleHandler.ListRoles)
			roles.GET("/:id", roleHandler.GetRole)
			roles.POST("", roleHandler.CreateRole)
			roles.PUT("/:id", roleHandler.UpdateRole)
			roles.DELETE("/:id", roleHandler.DeleteRole)
		}

		// Permission routes
		permissions := api.Group("/permissions")
		{
			permissions.GET("", permissionHandler.ListPermissions)
			permissions.GET("/:id", permissionHandler.GetPermission)
			permissions.POST("", permissionHandler.CreatePermission)
			permissions.PUT("/:id", permissionHandler.UpdatePermission)
			permissions.DELETE("/:id", permissionHandler.DeletePermission)
		}
	}
}
