package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Success(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func Error(message string) APIResponse {
	return APIResponse{
		Success: false,
		Message: message,
	}
}

func StatusOK() int {
	return http.StatusOK
}

// RespondSuccess writes a success response to the provided gin.Context
func RespondSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Success(message, data))
}

// RespondErrorStatus writes an error response with the given HTTP status
func RespondErrorStatus(c *gin.Context, status int, message string) {
	c.JSON(status, Error(message))
}
