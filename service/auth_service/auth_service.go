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
	ur user_repository.UserRepository
	pr photo_repository.PhotoRepository
	cr comment_repository.CommentRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository, commentRepo comment_repository.CommentRepository) AuthService {
	return &authServiceImpl{
		ur: userRepo,
		pr: photoRepo,
		cr: commentRepo,
	}
}

// Authentication implements a.
func (a *authServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		invalidToken := errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		user := entity.User{}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.ur.FetchById(user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidToken.Status(), invalidToken)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}

// AuthorizationPhoto implements a.
func (a *authServiceImpl) AuthorizationPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		photoId, _ := strconv.Atoi(ctx.Param("photoId"))

		photo, err := a.pr.GetPhotoId(photoId)

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

// AuthorizationComment implements a.
func (a *authServiceImpl) AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		commentId, _ := strconv.Atoi(ctx.Param("commentId"))

		comment, err := a.cr.GetCommentById(commentId)

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
