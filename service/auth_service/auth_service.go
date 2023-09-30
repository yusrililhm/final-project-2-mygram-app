package auth_service

import "github.com/gin-gonic/gin"

type AuthService interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
}

type authServiceImpl struct {
}

func NewAuthService() AuthService {
	return &authServiceImpl{}
}

// Authentication implements AuthService.
func (authService *authServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

// Authorization implements AuthService.
func (authService *authServiceImpl) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
