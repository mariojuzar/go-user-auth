package dao

import (
	"context"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	"github.com/mariojuzar/go-user-auth/internal/domain/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RolePermissionRepository struct {
	collection *mongo.Collection
}

func NewRolePermissionRepository(db *mongo.Database) repository.RolePermissionRepository {
	return &RolePermissionRepository{collection: db.Collection("role_permission")}
}

func (repo *RolePermissionRepository) InsertRolePermission(ctx context.Context, rp model.RolePermission) error {
	_, err := repo.collection.InsertOne(ctx, rp)
	if err != nil {
		logrus.WithError(err).Error("Failed to insert role permission")
		return err
	}
	return nil
}

func (repo *RolePermissionRepository) GetPermissionByRole(ctx context.Context, role string) ([]string, error) {
	var result []model.RolePermission
	filter := bson.D{
		{"role", role},
	}
	project := bson.D{
		{"permission", 1},
	}

	opts := options.Find().SetProjection(project)
	res, err := repo.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	err = res.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	var permissions []string
	for _, permission := range result {
		permissions = append(permissions, permission.Permission)
	}

	return permissions, err
}

func (repo *RolePermissionRepository) GetMapAccessPermissionByRole(ctx context.Context, role string) (map[string]string, error) {
	var result []model.RolePermission

	filter := bson.D{
		{"role", role},
	}
	project := bson.D{
		{"permission", 1},
		{"access_api", 1},
		{"method", 1},
	}

	opts := options.Find().SetProjection(project)
	res, err := repo.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	err = res.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	access := make(map[string]string, 0)
	for _, permission := range result {
		access[permission.AccessAPI+":"+permission.Method] = permission.Permission
	}

	return access, nil
}
