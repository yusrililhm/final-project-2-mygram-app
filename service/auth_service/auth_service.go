package auth_service

import (
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
}

type authServiceImpl struct {
	userRepository user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) AuthService {
	return &authServiceImpl{
		userRepository: userRepo,
	}
}

// Authentication implements AuthService.
func (authService *authServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		invalidToken := errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		user := entity.User{}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = authService.userRepository.FetchById(user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidToken.Status(), invalidToken)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}

// Authorization implements AuthService.
func (authService *authServiceImpl) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
