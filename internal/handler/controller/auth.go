package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mariojuzar/go-user-auth/internal/handler/request"
	"github.com/mariojuzar/go-user-auth/internal/handler/response"
	"github.com/mariojuzar/go-user-auth/pkg/custerr"
)

// LoginUser godoc
// @Summary		LoginUser
// @Description LoginUser
// @Tags 		auth
// @Accept		json
// @Param		user	body	request.LoginRequest 	true 	"login user payload"
// @Produce 	json
// @Success		200	{object} 	response.BaseResponse{data=response.AuthResponse}
// @Router		/v1/auth/login	[post]
func (a *API) LoginUser(ctx *fiber.Ctx) error {
	var req request.LoginRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}
	result, err := a.authUc.Login(ctx.Context(), req)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}

	return ctx.JSON(response.BaseResponse{
		Data: result,
	})
}

// RefreshToken godoc
// @Summary		RefreshToken
// @Description RefreshToken
// @Tags 		auth
// @Accept		json
// @Param		user	body	request.RefreshTokenRequest 	true 	"refresh token payload"
// @Produce 	json
// @Success		200	{object} 	response.BaseResponse{data=response.AuthResponse}
// @Router		/v1/auth/refresh	[post]
func (a *API) RefreshToken(ctx *fiber.Ctx) error {
	var req request.RefreshTokenRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}

	result, err := a.authUc.RefreshToken(ctx.Context(), req)
	if err != nil {
		return &custerr.CustomError{
			Err:     err,
			ErrCode: 400,
		}
	}

	return ctx.JSON(response.BaseResponse{
		Data: result,
	})
}
