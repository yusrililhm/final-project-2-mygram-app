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
	us user_service.UserService
}

func NewUserHandler(userService user_service.UserService) UserHandler {
	return &userHandlerImpl{
		us: userService,
	}
}
// Register implements UserHandler
// Register godoc
// @Summary User register
// @Description User register
// @Tags Users
// @Accept json
// @Produce json
// @Param dto.NewUserRequest body dto.NewUserRequest true "body request for user register"
// @Success 201 {object} dto.GetUserResponse
// @Router /users/register [post]
func (u *userHandlerImpl) Register(ctx *gin.Context) {
	userPayload := &dto.NewUserRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := u.us.Add(userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Login implements UserHandler.
// Login godoc
// @Summary User login
// @Description User login
// @Tags Users
// @Accept json
// @Produce json
// @Param dto.UserLoginRequest body dto.UserLoginRequest true "body request for user login"
// @Success 200 {object} dto.GetUserResponse
// @Router /users/login [post]
func (u *userHandlerImpl) Login(ctx *gin.Context) {
	userPayload := &dto.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := u.us.Get(userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Update implements UserHandler.
// Update godoc
// @Summary User update
// @Description User update
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param dto.UserUpdateRequest body dto.UserUpdateRequest true "body request for user update"
// @Success 200 {object} dto.GetUserResponse
// @Router /users [put]
func (u *userHandlerImpl) Update(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	userPayload := &dto.UserUpdateRequest{}

	if err := ctx.ShouldBindJSON(userPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := u.us.Edit(user.Id, userPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Delete implements UserHandler.
// Delete godoc
// @Summary Create new User
// @Description Create new Users
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.GetUserResponse
// @Router /users [delete]
func (u *userHandlerImpl) Delete(ctx *gin.Context) {

	user := ctx.MustGet("userData").(entity.User)

	response, err := u.us.Remove(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
