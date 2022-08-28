package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	"github.com/mariojuzar/go-user-auth/internal/handler/request"
	"github.com/mariojuzar/go-user-auth/internal/handler/response"
	"github.com/mariojuzar/go-user-auth/internal/interfaces/dao"
	"github.com/mariojuzar/go-user-auth/pkg/constant"
	"github.com/mariojuzar/go-user-auth/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (m *Module) FindById(ctx context.Context, userId primitive.ObjectID) (*response.UserResponse, error) {
	user, err := m.userRepo.FindById(ctx, userId)
	if err != nil {
		if err == dao.ErrNotFound {
			return nil, &fiber.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			}
		}
		return nil, &fiber.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return &response.UserResponse{
		UserId:    user.ID.Hex(),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      string(user.Role),
	}, nil
}

func (m *Module) DeleteById(ctx context.Context, userId primitive.ObjectID) error {
	user, err := m.userRepo.FindById(ctx, userId)
	if err != nil {
		if err == dao.ErrNotFound {
			return &fiber.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			}
		}
		return &fiber.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	err = m.userRepo.Delete(ctx, user.ID)
	if err != nil {
		return &fiber.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func (m *Module) CreateUser(ctx context.Context, req request.UserCreateRequest) error {
	jwt, err := utils.ReadJwtData(ctx)
	if err != nil {
		return err
	}

	role := model.UserRoles[req.Role]
	if role == constant.EmptyString {
		return &fiber.Error{
			Message: "invalid role",
			Code:    http.StatusBadRequest,
		}
	}

	rolePermission, err := m.rolePermissionRepo.GetPermissionByRole(ctx, req.Role)
	if err != nil {
		return &fiber.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	// encrypt password
	passwordBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 5)
	user := &model.User{
		Username:    req.Username,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Password:    string(passwordBytes),
		Role:        role,
		Permissions: rolePermission,
		CreatedAt:   time.Now(),
		CreatedBy:   jwt.UserId,
		UpdatedAt:   time.Now(),
		UpdatedBy:   jwt.UserId,
	}

	err = m.userRepo.Create(ctx, user)
	if err != nil {
		return &fiber.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func (m *Module) FindAll(ctx context.Context, page int, size int) ([]response.UserResponse, error) {
	res, err := m.userRepo.FindAll(ctx, page, size)
	if err != nil {
		return nil, &fiber.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	resp := make([]response.UserResponse, 0)
	for _, user := range res {
		resp = append(resp, response.UserResponse{
			UserId:    user.ID.Hex(),
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      string(user.Role),
		})
	}
	return resp, nil
}

func (m *Module) Update(ctx context.Context, req request.UserUpdateRequest) error {
	user, err := m.userRepo.FindById(ctx, req.UserId)
	if err != nil {
		if err == dao.ErrNotFound {
			return &fiber.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			}
		}
		return &fiber.Error{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	userUpdate := &model.UserUpdate{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if req.Username != constant.EmptyString {
		usernameCheck, err := m.userRepo.FindByUsername(ctx, req.Username)
		if err != nil && err != dao.ErrNotFound {
			return &fiber.Error{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
		}
		if usernameCheck != nil {
			return &fiber.Error{
				Message: "username already in use",
				Code:    http.StatusBadRequest,
			}
		}
		userUpdate.Username = req.Username
	}

	if req.Role != constant.EmptyString && string(user.Role) != req.Role {
		role := model.UserRoles[req.Role]
		if role == constant.EmptyString {
			return &fiber.Error{
				Message: "invalid role",
				Code:    http.StatusBadRequest,
			}
		}

		permission, err := m.rolePermissionRepo.GetPermissionByRole(ctx, req.Role)
		if err != nil {
			return &fiber.Error{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
		}

		userUpdate.Role = role
		userUpdate.Permissions = permission
	}

	if req.Password != constant.EmptyString {
		passwordBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 5)
		userUpdate.Password = string(passwordBytes)
	}

	err = m.userRepo.UpdateById(ctx, userUpdate, req.UserId)
	if err != nil {
		return err
	}
	return nil
}
