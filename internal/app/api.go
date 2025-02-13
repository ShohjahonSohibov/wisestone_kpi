package app

import (
	"kpi/config"
	"kpi/internal/handlers"
	"kpi/internal/middleware"
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
	roleHandler := handlers.NewRoleHandler(serviceManager.RoleService)
	permissionHandler := handlers.NewPermissionHandler(serviceManager.PermissionService)
	rolePermissionHandler := handlers.NewRolePermissionHandler(serviceManager.RolePermissionService)

	// API routes
	api := router.Group("/api/v1")
	{
		// Public routes (no auth required)
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes (auth required)
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(config.Load().Secret))
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("", userHandler.ListUsers)
				users.GET("/:email", userHandler.GetUser)
				users.POST("", userHandler.CreateUser)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
			}

			// Team routes
			teams := protected.Group("/teams")
			{
				teams.GET("", teamHandler.ListTeams)
				teams.GET("/:id", teamHandler.GetTeam)
				teams.POST("", teamHandler.CreateTeam)
				teams.PUT("/:id", teamHandler.UpdateTeam)
				teams.DELETE("/:id", teamHandler.DeleteTeam)
			}

			// Role routes
			roles := protected.Group("/roles")
			{
				roles.GET("", roleHandler.ListRoles)
				roles.GET("/:id", roleHandler.GetRole)
				roles.POST("", roleHandler.CreateRole)
				roles.PUT("/:id", roleHandler.UpdateRole)
				roles.DELETE("/:id", roleHandler.DeleteRole)
			}

			// Permission routes
			permissions := protected.Group("/permissions")
			{
				permissions.GET("", permissionHandler.ListPermissions)
				permissions.GET("/:id", permissionHandler.GetPermission)
				permissions.POST("", permissionHandler.CreatePermission)
				permissions.PUT("/:id", permissionHandler.UpdatePermission)
				permissions.DELETE("/:id", permissionHandler.DeletePermission)
			}

			// Role-Permission routes
			rolePermissions := protected.Group("/role-permissions")
			{
				rolePermissions.GET("", rolePermissionHandler.ListRolePermissions)
				rolePermissions.GET("/:id", rolePermissionHandler.GetRolePermission)
				rolePermissions.POST("", rolePermissionHandler.CreateRolePermission)
				rolePermissions.PUT("/:id", rolePermissionHandler.UpdateRolePermission)
				rolePermissions.DELETE("/:id", rolePermissionHandler.DeleteRolePermission)
			}
		}
	}
}
