package usecases

import (
	"context"
	"github.com/mariojuzar/go-user-auth/internal/handler/request"
	"github.com/mariojuzar/go-user-auth/internal/handler/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCase interface {
	FindById(ctx context.Context, userId primitive.ObjectID) (*response.UserResponse, error)
	DeleteById(ctx context.Context, userId primitive.ObjectID) error
	CreateUser(ctx context.Context, req request.UserCreateRequest) error
	FindAll(ctx context.Context, page int, size int) ([]response.UserResponse, error)
	Update(ctx context.Context, req request.UserUpdateRequest) error
}
