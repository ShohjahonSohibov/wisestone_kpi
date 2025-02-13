package handlers

import (
	"kpi/internal/models"
	"kpi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login godoc
// @Summary Authenticate an existing user
// @Description Login with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login Request"
// @Success 200 {object} models.LoginResponse "token: JWT Token"
// @Failure 401 {object} map[string]string "error: Unauthorized"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    req := &models.LoginRequest{}

    if err := c.ShouldBindJSON(req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": err.Error(),
        })
        return
    }

    // Validate required fields
    if req.Username == "" || req.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "username and password are required",
        })
        return
    }

    res, err := h.authService.Login(c.Request.Context(), req.Username, req.Password)
    if err != nil {
        // Handle different types of errors with appropriate status codes
        switch err.Error() {
        case "user not found", "invalid email or password":
            c.JSON(http.StatusUnauthorized, gin.H{
                "status":  http.StatusUnauthorized,
                "message": err.Error(),
            })
        default:
            c.JSON(http.StatusInternalServerError, gin.H{
                "status":  http.StatusInternalServerError,
                "message": "internal server error",
            })
        }
        return
    }

    if res == nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "failed to generate login response",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   res,
    })
}
