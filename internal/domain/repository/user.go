package repository

import (
	"context"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	FindById(ctx context.Context, userId primitive.ObjectID) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	UpdateById(ctx context.Context, user *model.UserUpdate, userId primitive.ObjectID) error
	Delete(ctx context.Context, userId primitive.ObjectID) error
	FindAll(ctx context.Context, page, size int) ([]model.User, error)
}
