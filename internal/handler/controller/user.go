package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mariojuzar/go-user-auth/internal/handler/request"
	"github.com/mariojuzar/go-user-auth/internal/handler/response"
	"github.com/mariojuzar/go-user-auth/pkg/custerr"
	"github.com/mariojuzar/go-user-auth/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

// FindUserById godoc
// @Summary		FindUserById
// @Description FindUserById
// @Tags 		users
// @Accept		json
// @Security 	ApiKeyAuth
// @Param		user_id		path	string		true 	"user id"
// @Produce 	json
// @Success		200	{object} 	response.BaseResponse{data=response.UserResponse}
// @Router		/v1/users/user/:user_id	[get]
func (a *API) FindUserById(ctx *fiber.Ctx) error {
	var id primitive.ObjectID
	userIdStr := ctx.Params("user_id")
	id, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}
	result, err := a.userUc.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{Data: result})
}

// FindMyUser godoc
// @Summary		FindMyUser
// @Description FindMyUser
// @Tags 		users
// @Accept		json
// @Security 	ApiKeyAuth
// @Produce 	json
// @Success		200	{object} 	response.BaseResponse{data=response.UserResponse}
// @Router		/v1/users/me	[get]
func (a *API) FindMyUser(ctx *fiber.Ctx) error {
	var id primitive.ObjectID

	jwtData, err := utils.ReadJwtData(ctx.Context())
	if err != nil {
		return err
	}
	id, err = primitive.ObjectIDFromHex(jwtData.UserId)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}
	result, err := a.userUc.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{Data: result})
}

// FindAllUser godoc
// @Summary		FindAllUser
// @Description FindAllUser
// @Tags 		users
// @Accept		json
// @Security 	ApiKeyAuth
// @Param		page	query	int 	false 	"page"
// @Param		size	query	int 	false 	"size result"
// @Produce 	json
// @Success		200	{object} 	response.BaseResponse{data=[]response.UserResponse}
// @Router		/v1/users	[get]
func (a *API) FindAllUser(ctx *fiber.Ctx) error {
	var page, size int
	p := ctx.Query("page")
	s := ctx.Query("size")
	page, err := strconv.Atoi(p)
	if err != nil || page < 1 {
		page = 0
	}
	size, err = strconv.Atoi(s)
	if err != nil || size < 1 {
		size = 10
	}

	result, err := a.userUc.FindAll(ctx.Context(), page, size)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{Data: result})
}

// CreateUser godoc
// @Summary		CreateUser
// @Description CreateUser
// @Tags 		users
// @Accept		json
// @Security 	ApiKeyAuth
// @Param		user	body	request.UserCreateRequest 	true 	"create user payload"
// @Produce 	json
// @Success		200	{object} 	response.BaseResponse{}
// @Router		/v1/users	[post]
func (a *API) CreateUser(ctx *fiber.Ctx) error {
	var req request.UserCreateRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}

	err = a.userUc.CreateUser(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{Message: "user created"})
}

// UpdateUser godoc
// @Summary		UpdateUser
// @Description UpdateUser
// @Tags 		users
// @Accept		json
// @Security 	ApiKeyAuth
// @Param		user	body	request.UserUpdateRequest 	true 	"update user payload"
// @Produce 	json
// @Success		200	{object} 	response.BaseResponse{}
// @Router		/v1/users	[patch]
func (a *API) UpdateUser(ctx *fiber.Ctx) error {
	var req request.UserUpdateRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}

	err = a.userUc.Update(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{Message: "user updated"})
}

// DeleteUser godoc
// @Summary		DeleteUser
// @Description DeleteUser
// @Tags 		users
// @Accept		json
// @Security 	ApiKeyAuth
// @Param		user_id		path	string		true 	"user id"
// @Produce 	json
// @Success		200	{object} 	response.BaseResponse{}
// @Router		/v1/users/:user_id	[delete]
func (a *API) DeleteUser(ctx *fiber.Ctx) error {
	var id primitive.ObjectID
	userIdStr := ctx.Params("user_id")
	id, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}

	err = a.userUc.DeleteById(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{Message: "user deleted"})
}
