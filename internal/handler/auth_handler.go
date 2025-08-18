package handler

import (
	"net/http"

	"github.com/Rambudhi/payslip/internal/request"
	"github.com/Rambudhi/payslip/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

// @Summary Login
// @Description Login with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login Request"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Get user profile
// @Description Get current logged-in user's info from JWT
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/me [get]
func (h *AuthHandler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user_id":   c.GetUint("user_id"),
		"user_name": c.GetString("user_name"),
		"role":      c.GetString("role"),
	})
}
