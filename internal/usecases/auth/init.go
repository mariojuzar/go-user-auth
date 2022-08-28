package auth

import (
	"github.com/mariojuzar/go-user-auth/internal/config"
	"github.com/mariojuzar/go-user-auth/internal/domain/repository"
	"github.com/mariojuzar/go-user-auth/internal/usecases"
)

type Opts struct {
	UserRepo   repository.UserRepository
	AuthConfig config.AuthConfig
}

type Module struct {
	userRepo   repository.UserRepository
	authConfig config.AuthConfig
}

func New(o *Opts) usecases.AuthUseCase {
	return &Module{
		userRepo:   o.UserRepo,
		authConfig: o.AuthConfig,
	}
}
