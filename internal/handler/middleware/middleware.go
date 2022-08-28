package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mariojuzar/go-user-auth/internal/config"
	"github.com/mariojuzar/go-user-auth/internal/domain/repository"
)

type Opts struct {
	CfgAuth            config.AuthConfig
	RolePermissionRepo repository.RolePermissionRepository
}

type interceptor struct {
	cfgAuth            config.AuthConfig
	rolePermissionRepo repository.RolePermissionRepository
}

type Interceptor interface {
	Auth(h fiber.Handler) fiber.Handler
}

func NewMiddleware(o *Opts) Interceptor {
	return &interceptor{
		cfgAuth:            o.CfgAuth,
		rolePermissionRepo: o.RolePermissionRepo,
	}
}
