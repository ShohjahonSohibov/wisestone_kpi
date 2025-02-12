package handlers

import (
	"kpi/internal/models"
	"kpi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RolePermissionHandler struct {
	rolePermissionService *services.RolePermissionService
}

func NewRolePermissionHandler(rolePermissionService *services.RolePermissionService) *RolePermissionHandler {
	return &RolePermissionHandler{rolePermissionService: rolePermissionService}
}

// GetRolePermission godoc
// @Summary Get role-permission by id
// @Description Get role-permission details by id
// @Tags RolePermissions
// @Accept json
// @Produce json
// @Param id path string true "RolePermission Id"
// @Success 200 {object} models.RolePermission
// @Failure 404 {object} map[string]string "error: RolePermission not found"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/role-permissions/{id} [get]
func (h *RolePermissionHandler) GetRolePermission(c *gin.Context) {
	id := c.Param("id")
	rolePermission, err := h.rolePermissionService.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rolePermission)
}

// CreateRolePermission godoc
// @Summary Create a new role-permission
// @Description Create a new role-permission with the provided details
// @Tags RolePermissions
// @Accept json
// @Produce json
// @Param rolePermission body models.CreateRolePermission true "RolePermission Details"
// @Success 201 {object} map[string]string "message: RolePermission created successfully"
// @Failure 400 {object} map[string]string "error: Validation error"
// @Router /api/v1/role-permissions [post]
func (h *RolePermissionHandler) CreateRolePermission(c *gin.Context) {
	var rolePermission models.RolePermission
	if err := c.ShouldBindJSON(&rolePermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.rolePermissionService.Create(c.Request.Context(), &rolePermission); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "role permission created successfully"})
}

// UpdateRolePermission godoc
// @Summary Update existing role-permission
// @Description Update role-permission details by ID
// @Tags RolePermissions
// @Accept json
// @Produce json
// @Param id path string true "RolePermission ID"
// @Param rolePermission body models.UpdateRolePermission true "RolePermission Details"
// @Success 200 {object} map[string]string "message: RolePermission updated successfully"
// @Failure 400 {object} map[string]string "error: Validation error"
// @Failure 404 {object} map[string]string "error: RolePermission not found"
// @Router /api/v1/role-permissions/{id} [put]
func (h *RolePermissionHandler) UpdateRolePermission(c *gin.Context) {
	id := c.Param("id")
	var rolePermission models.UpdateRolePermission
	if err := c.ShouldBindJSON(&rolePermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rolePermission.ID = id
	if err := h.rolePermissionService.Update(c.Request.Context(), &rolePermission); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role permission updated successfully"})
}

// DeleteRolePermission godoc
// @Summary Delete role-permission
// @Description Delete role-permission by ID
// @Tags RolePermissions
// @Accept json
// @Produce json
// @Param id path string true "RolePermission Id"
// @Success 200 {object} map[string]string "message: RolePermission deleted successfully"
// @Failure 404 {object} map[string]string "error: RolePermission not found"
// @Router /api/v1/role-permissions/{id} [delete]
func (h *RolePermissionHandler) DeleteRolePermission(c *gin.Context) {
	id := c.Param("id")

	if err := h.rolePermissionService.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "role permission deleted successfully"})
}

// ListRolePermissions godoc
// @Summary List all role-permissions
// @Description Get a list of all role-permissions
// @Tags RolePermissions
// @Accept json
// @Produce json
// @Param filter query models.ListRolePermissionRequest false "Filter parameters"
// @Success 200 {object} models.ListRolePermissionResponse
// @Failure 400 {object} map[string]string "error: Invalid parameters"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/role-permissions [get]
func (h *RolePermissionHandler) ListRolePermissions(c *gin.Context) {
	filter := &models.ListRolePermissionRequest{}

	offset, limit, err := getPageOffsetLimit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter.Filter.Offset = offset
	filter.Filter.Limit = limit
	filter.RoleId = c.Query("role_id")
	filter.PermissionId = c.Query("permission_id")

	rolePermissions, err := h.rolePermissionService.List(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rolePermissions)
}
