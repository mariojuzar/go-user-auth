package utils

import (
	"context"
	"fmt"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	"github.com/mariojuzar/go-user-auth/pkg/custerr"
)

func ReadJwtData(ctx context.Context) (*model.JwtData, error) {
	val := ctx.Value(model.JwtKey)
	jwt, ok := val.(*model.JwtClaim)
	if !ok {
		return nil, &custerr.CustomError{
			Err:     fmt.Errorf("unauthorized"),
			ErrCode: 401,
		}
	}

	return &jwt.JwtData, nil
}
