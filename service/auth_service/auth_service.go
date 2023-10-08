package auth_service

import (
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/comment_repository"
	"myGram/repository/photo_repository"
	"myGram/repository/user_repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AuthorizationPhoto() gin.HandlerFunc
	AuthorizationComment() gin.HandlerFunc
}

type authServiceImpl struct {
	userRepository    user_repository.UserRepository
	photoRepository   photo_repository.PhotoRepository
	commentRepository comment_repository.CommentRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository, commentRepo comment_repository.CommentRepository) AuthService {
	return &authServiceImpl{
		userRepository:    userRepo,
		photoRepository:   photoRepo,
		commentRepository: commentRepo,
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

// AuthorizationPhoto implements AuthService.
func (authService *authServiceImpl) AuthorizationPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		photoId, _ := strconv.Atoi(ctx.Param("photoId"))

		photo, err := authService.photoRepository.GetPhotoId(photoId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserId != user.Id {
			errUnathorized := errs.NewUnathorizedError("you are not authorized to modify the photo")
			ctx.AbortWithStatusJSON(errUnathorized.Status(), errUnathorized)
		}

		ctx.Next()
	}
}

// AuthorizationComment implements AuthService.
func (authService *authServiceImpl) AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		commentId, _ := strconv.Atoi(ctx.Param("commentId"))

		comment, err := authService.commentRepository.GetCommentById(commentId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if comment.UserId != user.Id {
			errUnathorized := errs.NewUnathorizedError("you are not authorized to modify the comment")
			ctx.AbortWithStatusJSON(errUnathorized.Status(), errUnathorized)
		}

		ctx.Next()
	}
}
