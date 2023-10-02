package handler

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/service/user_service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userHandlerImpl struct {
	userService user_service.UserService
}

func NewUserHandler(userService user_service.UserService) UserHandler {
	return &userHandlerImpl{
		userService: userService,
	}
}

func (userHandler *userHandlerImpl) Register(ctx *gin.Context) {
	userPayload := &dto.NewUserRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := userHandler.userService.Add(userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Login implements UserHandler.
func (userHandler *userHandlerImpl) Login(ctx *gin.Context) {
	userPayload := &dto.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := userHandler.userService.Get(userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Update implements UserHandler.
func (userHandler *userHandlerImpl) Update(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	userPayload := &dto.UserUpdateRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := userHandler.userService.Edit(user.Id, userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Delete implements UserHandler.
func (userHandler *userHandlerImpl) Delete(ctx *gin.Context) {

	user := ctx.MustGet("userData").(entity.User)

	response, err := userHandler.userService.Remove(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
