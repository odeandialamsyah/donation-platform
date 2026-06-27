package response

import "net/http"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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
