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
}

type userHandlerImpl struct {
	userService user_service.UserService
}

func NewUserHandler(userService user_service.UserService) UserHandler {
	return &userHandlerImpl{
		userService: userService,
	}
}

func (uh *userHandlerImpl) Register(ctx *gin.Context) {
	userPayload := &dto.NewUserRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Add(userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Login implements UserHandler.
func (uh *userHandlerImpl) Login(ctx *gin.Context) {
	userPayload := &dto.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request " + err.Error())
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Get(userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Update implements UserHandler.
func (uh *userHandlerImpl) Update(ctx *gin.Context) {
	userId := ctx.MustGet("userData").(entity.User)

	userPayload := &dto.UserUpdateRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request " + err.Error())
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Edit(userId.Id, userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
