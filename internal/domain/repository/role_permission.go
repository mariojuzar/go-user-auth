package repository

import (
	"context"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
)

type RolePermissionRepository interface {
	InsertRolePermission(ctx context.Context, rp model.RolePermission) error
	GetPermissionByRole(ctx context.Context, role string) ([]string, error)
	GetMapAccessPermissionByRole(ctx context.Context, role string) (map[string]string, error)
}
