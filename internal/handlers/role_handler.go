package handlers

import (
	"kpi/internal/models"
	"kpi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

// GetRole godoc
// @Security ApiKeyAuth
// @Summary Get role by id
// @Description Get role details by id
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role Id"
// @Success 200 {object} map[string]interface{} "status: 200, data: Role"
// @Failure 401 {object} map[string]interface{} "status: 401, message: Unauthorized"
// @Failure 404 {object} map[string]interface{} "status: 404, message: Role not found"
// @Router /api/v1/roles/{id} [get]
func (h *RoleHandler) GetRole(c *gin.Context) {
	id := c.Param("id")
	role, err := h.roleService.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   role,
	})
}

// CreateRole godoc
// @Security ApiKeyAuth
// @Summary Create a new role
// @Description Create a new role with the provided details
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body models.Role true "Role Details"
// @Success 201 {object} map[string]interface{} "status: 201, message: role created successfully"
// @Failure 401 {object} map[string]interface{} "status: 401, message: Unauthorized"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 409 {object} map[string]interface{} "status: 409, message: error message"
// @Router /api/v1/roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if err := h.roleService.Create(c.Request.Context(), &role); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "role created successfully",
	})
}

// UpdateRole godoc
// @Security ApiKeyAuth
// @Summary Update existing role
// @Description Update role details by ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param role body models.Role true "Role Details"
// @Success 200 {object} map[string]interface{} "status: 200, message: role updated successfully"
// @Failure 401 {object} map[string]interface{} "status: 401, message: Unauthorized"
// @Failure 400 {object} map[string]interface{} "status: 400, message: error message"
// @Failure 404 {object} map[string]interface{} "status: 404, message: Role not found"
// @Router /api/v1/roles/{id} [put]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	role.ID = id
	if err := h.roleService.Update(c.Request.Context(), &role); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "role updated successfully",
	})
}

// DeleteRole godoc
// @Security ApiKeyAuth
// @Summary Delete role
// @Description Delete role by ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role Id"
// @Success 200 {object} map[string]interface{} "status: 200, message: role deleted successfully"
// @Failure 401 {object} map[string]interface{} "status: 401, message: Unauthorized"
// @Failure 404 {object} map[string]interface{} "status: 404, message: Role not found"
// @Router /api/v1/roles/{id} [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if len(id) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid ID format",
		})
		return
	}

	if err := h.roleService.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "role deleted successfully",
	})
}

// ListRoles godoc
// @Security ApiKeyAuth
// @Summary List all roles
// @Description Get a list of all roles
// @Tags Roles
// @Accept json
// @Produce json
// @Param filter query models.ListRoleRequest false "Filter parameters"
// @Success 200 {object} map[string]interface{} "status: 200, data: roles"
// @Failure 401 {object} map[string]interface{} "status: 401, message: Unauthorized"
// @Failure 400 {object} map[string]interface{} "status: 400, message: Invalid parameters"
// @Failure 500 {object} map[string]interface{} "status: 500, message: Internal server error"
// @Router /api/v1/roles [get]
func (h *RoleHandler) ListRoles(c *gin.Context) {
	filter := &models.ListRoleRequest{}
	if err := c.ShouldBindQuery(filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

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

	roles, err := h.roleService.List(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   roles,
	})
}
