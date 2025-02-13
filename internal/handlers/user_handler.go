package handlers

import (
	"kpi/internal/models"
	"kpi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUser godoc
// @Summary Get user by id
// @Description Get user details by id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string "error: User not found"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "User Details"
// @Success 201 {object} map[string]string "message: User created successfully"
// @Failure 400 {object} map[string]string "error: Validation error"
// @Failure 409 {object} map[string]string "error: User already exists"
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if err := h.userService.Create(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "user created successfully",
	})
}

// UpdateUser godoc
// @Summary Update existing user
// @Description Update user details by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "User Details"
// @Success 200 {object} map[string]string "message: User updated successfully"
// @Failure 400 {object} map[string]string "error: Validation error"
// @Failure 404 {object} map[string]string "error: User not found"
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user.ID = id
	if err := h.userService.Update(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user updated successfully",
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by email
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Success 200 {object} map[string]string "message: User deleted successfully"
// @Failure 404 {object} map[string]string "error: User not found"
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if len(id) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid ID format",
		})
		return
	}

	if err := h.userService.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user deleted successfully",
	})
}

// ListUsers godoc
// @Summary List all users
// @Description Get a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Param filter query models.ListUsersRequest false "Filter parameters"
// @Success 200 {object} models.ListUsersResponse
// @Failure 400 {object} map[string]string "error: Invalid parameters"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	filter := &models.ListUsersRequest{}

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

	users, err := h.userService.List(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
	})
}
