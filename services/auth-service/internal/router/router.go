package router

import (
	"donation-platform/auth-service/internal/auth"
	"donation-platform/auth-service/internal/config"
	"donation-platform/auth-service/internal/handler"
	"donation-platform/auth-service/internal/repository"
	"donation-platform/auth-service/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {

	r := gin.Default()

	healthHandler := handler.NewHealthHandler()

	userRepo := repository.NewUserRepository(db)
	jwtSvc := &auth.JWTService{Secret: cfg.JWTSecret}
	authService := service.NewAuthService(userRepo, jwtSvc)
	authHandler := handler.NewAuthHandler(authService)

	r.GET("/health", healthHandler.Health)
	r.POST("/api/auth/register", authHandler.Register)

	return r
}
