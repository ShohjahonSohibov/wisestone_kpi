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
	kpiParentHandler := handlers.NewKPIParentHandler(serviceManager.KPIParentService)
	kpiDivisionHandler := handlers.NewKPIDivisionHandler(serviceManager.KPIDivisionService)
	kpiCriterionHandler := handlers.NewKPICriterionHandler(serviceManager.KPICriterionService)
	kpiFactorHandler := handlers.NewKPIFactorHandler(serviceManager.KPIFactorService)
	kpiFactorIndicatorHandler := handlers.NewKPIFactorIndicatorHandler(serviceManager.KPIFactorIndicatorService)
	kpiProgressHandler := handlers.NewKPIProgressHandler(serviceManager.KPIProgressService)
	kpiProgressStatusHandler := handlers.NewKPIProgressStatusHandler(serviceManager.KPIProgressStatusService)

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
				users.GET("/single", userHandler.GetUser)
				users.POST("", userHandler.CreateUser)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
			}

			// User team management routes (separate group to avoid conflicts)
			userTeams := protected.Group("/user-teams")
			{
				userTeams.PUT("", userHandler.AssignTeam)
				userTeams.DELETE("/:user_id", userHandler.RemoveFromTeam)
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

			// KPI-Parent routes
			kpiParentGroup := protected.Group("/kpi-parents")
			{
				kpiParentGroup.POST("", kpiParentHandler.Create)
				kpiParentGroup.PUT("/:id", kpiParentHandler.Update)
				kpiParentGroup.PUT("/status/:id", kpiParentHandler.Update)
				kpiParentGroup.DELETE("/:id", kpiParentHandler.Delete)
				kpiParentGroup.GET("/single", kpiParentHandler.GetByID)
				kpiParentGroup.GET("", kpiParentHandler.List)
			}

			// Kpi-Division routes
			kpiDivisionGroup := protected.Group("/kpi-divisions")
			{
				kpiDivisionGroup.POST("", kpiDivisionHandler.Create)
				kpiDivisionGroup.PUT("/:id", kpiDivisionHandler.Update)
				kpiDivisionGroup.DELETE("/:id", kpiDivisionHandler.Delete)
				kpiDivisionGroup.GET("/:id", kpiDivisionHandler.GetByID)
				kpiDivisionGroup.GET("", kpiDivisionHandler.List)
			}

			// Kpi-Criterion routes
			kpiCriterionGroup := protected.Group("/kpi-criterions")
			{
				kpiCriterionGroup.POST("", kpiCriterionHandler.Create)
				kpiCriterionGroup.PUT("/:id", kpiCriterionHandler.Update)
				kpiCriterionGroup.DELETE("/:id", kpiCriterionHandler.Delete)
				kpiCriterionGroup.GET("/:id", kpiCriterionHandler.GetByID)
				kpiCriterionGroup.GET("", kpiCriterionHandler.List)
			}

			// Kpi-Factor routes
			kpiFactorGroup := protected.Group("/kpi-factors")
			{
				kpiFactorGroup.POST("", kpiFactorHandler.Create)
				kpiFactorGroup.PUT("/:id", kpiFactorHandler.Update)
				kpiFactorGroup.DELETE("/:id", kpiFactorHandler.Delete)
				kpiFactorGroup.GET("/:id", kpiFactorHandler.GetByID)
				kpiFactorGroup.GET("", kpiFactorHandler.List)
			}

			// Kpi-Factor-Indicator routes
			kpiFactorIndicatorGroup := protected.Group("/kpi-factor-indicators")
			{
				kpiFactorIndicatorGroup.POST("", kpiFactorIndicatorHandler.Create)
				kpiFactorIndicatorGroup.PUT("/:id", kpiFactorIndicatorHandler.Update)
				kpiFactorIndicatorGroup.DELETE("/:id", kpiFactorIndicatorHandler.Delete)
				kpiFactorIndicatorGroup.GET("/:id", kpiFactorIndicatorHandler.GetByID)
				kpiFactorIndicatorGroup.GET("", kpiFactorIndicatorHandler.List)
			}

			kpiProgress := protected.Group("/kpi-progresses")
			{
				kpiProgress.POST("", kpiProgressHandler.Create)
				kpiProgress.POST("/bulk", kpiProgressHandler.CreateMany)
				kpiProgress.GET("", kpiProgressHandler.List)
				// kpiProgress.GET("/:id", kpiProgressHandler.GetByID)
				// kpiProgress.PUT("/:id", kpiProgressHandler.Update)
				kpiProgress.DELETE("/delete", kpiProgressHandler.Delete)
			}

			kpiProgressStatus := protected.Group("/kpi-progress-status")
			{
				kpiProgressStatus.POST("", kpiProgressStatusHandler.Create)
				kpiProgressStatus.GET("", kpiProgressStatusHandler.List)
				kpiProgressStatus.PUT("/:id", kpiProgressStatusHandler.Update)
				kpiProgressStatus.DELETE("/:id", kpiProgressStatusHandler.Delete)
			}
		}
	}
}
