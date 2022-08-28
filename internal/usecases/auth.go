package usecases

import (
	"context"
	"github.com/mariojuzar/go-user-auth/internal/handler/request"
	"github.com/mariojuzar/go-user-auth/internal/handler/response"
)

type AuthUseCase interface {
	Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error)
	RefreshToken(ctx context.Context, req request.RefreshTokenRequest) (*response.AuthResponse, error)
}
