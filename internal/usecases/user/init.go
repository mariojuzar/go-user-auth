package user

import (
	"github.com/mariojuzar/go-user-auth/internal/domain/repository"
	"github.com/mariojuzar/go-user-auth/internal/usecases"
)

type Opts struct {
	UserRepo           repository.UserRepository
	RolePermissionRepo repository.RolePermissionRepository
}

type Module struct {
	userRepo           repository.UserRepository
	rolePermissionRepo repository.RolePermissionRepository
}

func New(o *Opts) usecases.UserUseCase {
	return &Module{
		userRepo:           o.UserRepo,
		rolePermissionRepo: o.RolePermissionRepo,
	}
}
