package handlers

import (
	"kpi/internal/models"
	"kpi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	permissionService *services.PermissionService
}

func NewPermissionHandler(permissionService *services.PermissionService) *PermissionHandler {
	return &PermissionHandler{permissionService: permissionService}
}

// GetPermission godoc
// @Security ApiKeyAuth
// @Summary Get permission by id
// @Description Get permission details by id
// @Tags Permissions
// @Accept json
// @Produce json
// @Param id path string true "Permission Id"
// @Success 200 {object} models.Permission
// @Failure 404 {object} map[string]string "error: Permission not found"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/permissions/{id} [get]
func (h *PermissionHandler) GetPermission(c *gin.Context) {
	id := c.Param("id")
	permission, err := h.permissionService.GetById(c.Request.Context(), id)
	if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": err.Error(),
			})
			return
	}
	c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   permission,
	})
}

// CreatePermission godoc
// @Security ApiKeyAuth
// @Summary Create a new permission
// @Description Create a new permission with the provided details
// @Tags Permissions
// @Accept json
// @Produce json
// @Param permission body models.CreatePermission true "Permission Details"
// @Success 201 {object} map[string]string "message: Permission created successfully"
// @Failure 400 {object} map[string]string "error: Validation error"
// @Router /api/v1/permissions [post]
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	permission := &models.Permission{}
	if err := c.ShouldBindJSON(&permission); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": err.Error(),
			})
			return
	}
	if err := h.permissionService.Create(c.Request.Context(), permission); err != nil {
			c.JSON(http.StatusConflict, gin.H{
					"status":  http.StatusConflict,
					"message": err.Error(),
			})
			return
	}

	c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "permission created successfully",
	})
}

// UpdatePermission godoc
// @Security ApiKeyAuth
// @Summary Update existing permission
// @Description Update permission details by ID
// @Tags Permissions
// @Accept json
// @Produce json
// @Param id path string true "Permission ID"
// @Param permission body models.UpdatePermission true "Permission Details"
// @Success 200 {object} map[string]string "message: Permission updated successfully"
// @Failure 400 {object} map[string]string "error: Validation error"
// @Failure 404 {object} map[string]string "error: Permission not found"
// @Router /api/v1/permissions/{id} [put]
func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	id := c.Param("id")
	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": err.Error(),
			})
			return
	}

	permission.ID = id
	if err := h.permissionService.Update(c.Request.Context(), &permission); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": err.Error(),
			})
			return
	}

	c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "permission updated successfully",
	})
}

// DeletePermission godoc
// @Security ApiKeyAuth
// @Summary Delete permission
// @Description Delete permission by ID
// @Tags Permissions
// @Accept json
// @Produce json
// @Param id path string true "Permission Id"
// @Success 200 {object} map[string]string "message: Permission deleted successfully"
// @Failure 404 {object} map[string]string "error: Permission not found"
// @Router /api/v1/permissions/{id} [delete]
func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	if err := h.permissionService.Delete(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": err.Error(),
			})
			return
	}
	c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "permission deleted successfully",
	})
}

// ListPermissions godoc
// @Security ApiKeyAuth
// @Summary List all permissions
// @Description Get a list of all permissions
// @Tags Permissions
// @Accept json
// @Produce json
// @Param filter query models.ListPermissionRequest false "Filter parameters"
// @Success 200 {object} models.ListPermissionResponse
// @Failure 400 {object} map[string]string "error: Invalid parameters"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/permissions [get]
func (h *PermissionHandler) ListPermissions(c *gin.Context) {
	filter := &models.ListPermissionRequest{}

	offset, limit, err := getPageOffsetLimit(c)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": err.Error(),
			})
			return
	}
	filter.Filter.Offset = offset
	filter.Filter.Limit = limit
	filter.MultiSearch = c.Query("multi_search")
	filter.SortOrder = c.Query("sort_order")

	permissions, err := h.permissionService.List(c.Request.Context(), filter)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
					"status":  http.StatusInternalServerError,
					"message": err.Error(),
			})
			return
	}
	c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   permissions,
	})
}
