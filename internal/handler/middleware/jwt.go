package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	"github.com/mariojuzar/go-user-auth/internal/handler/response"
	"github.com/mariojuzar/go-user-auth/pkg/constant"
	"github.com/mariojuzar/go-user-auth/pkg/utils"
	"net/http"
	"strings"
)

func (i *interceptor) Auth(h fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		reqToken := ctx.GetReqHeaders()["Authorization"]
		if reqToken == constant.EmptyString {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(response.BaseResponse{
				Message: "unauthorized",
			})
		}

		splitToken := strings.Split(reqToken, "Bearer ")
		bearerToken := splitToken[1]
		if bearerToken == constant.EmptyString {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(response.BaseResponse{
				Message: "unauthorized",
			})
		}

		var claims model.JwtClaim
		token, err := jwt.ParseWithClaims(bearerToken, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, &fiber.Error{
					Code:    http.StatusBadRequest,
					Message: "invalid refresh token",
				}
			}

			return []byte(i.cfgAuth.JwtSecret), nil
		})
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(response.BaseResponse{
				Message: "unauthorized",
			})
		}

		claim, ok := token.Claims.(*model.JwtClaim)
		if ok && token.Valid {
			ctx.Locals(model.JwtKey, claim)
		} else {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(response.BaseResponse{
				Message: "unauthorized",
			})
		}

		if err = i.validatePermission(ctx, claim.JwtData); err != nil {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(response.BaseResponse{
				Message: "unauthorized",
			})
		}

		return h(ctx)
	}
}

func (i *interceptor) validatePermission(ctx *fiber.Ctx, jwtData model.JwtData) error {
	if jwtData.Role == model.SuperAdmin {
		return nil
	}

	if len(jwtData.Permissions) == 0 {
		return fmt.Errorf("user have no permissions")
	}

	access, err := i.rolePermissionRepo.GetMapAccessPermissionByRole(ctx.Context(), string(jwtData.Role))
	if err != nil {
		return err
	}

	accessPermission := ctx.Route().Path + ":" + ctx.Route().Method
	permission := access[accessPermission]

	if permission == constant.EmptyString {
		return fmt.Errorf("user have no permissions")
	}

	isExist, _ := utils.IsItemExist(jwtData.Permissions, permission)
	if !isExist {
		return fmt.Errorf("user have no permissions")
	}

	return nil
}
