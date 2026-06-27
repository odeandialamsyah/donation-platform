package handler

import (
	"donation-platform/auth-service/internal/request"
	"donation-platform/auth-service/internal/service"
	"donation-platform/shared/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondErrorStatus(c, 400, err.Error())
		return
	}

	accessToken, refreshToken, err := h.Service.Register(req.Name, req.Email, req.Password)

	if err != nil {
		response.RespondErrorStatus(c, 500, err.Error())
		return
	}

	response.RespondSuccess(c, "Register success", response.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.RespondErrorStatus(c, 400, err.Error())
		return
	}

	accessToken, refreshToken, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		response.RespondErrorStatus(c, 401, err.Error())
		return
	}

	response.RespondSuccess(c, "Login success", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
