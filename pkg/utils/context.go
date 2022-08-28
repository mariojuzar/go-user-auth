package utils

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	"net/http"
)

func ReadJwtData(ctx context.Context) (*model.JwtData, error) {
	val := ctx.Value(model.JwtKey)
	jwt, ok := val.(*model.JwtClaim)
	if !ok {
		return nil, &fiber.Error{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	return &jwt.JwtData, nil
}
