package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mariojuzar/go-user-auth/internal/domain/model"
	"github.com/mariojuzar/go-user-auth/internal/handler/request"
	"github.com/mariojuzar/go-user-auth/internal/handler/response"
	"github.com/mariojuzar/go-user-auth/internal/interfaces/dao"
	"github.com/mariojuzar/go-user-auth/pkg/custerr"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (m Module) Login(ctx context.Context, req request.LoginRequest) (*response.AuthResponse, error) {
	user, err := m.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		if err == dao.ErrNotFound {
			return nil, &custerr.CustomError{
				Err:     err,
				ErrCode: http.StatusBadRequest,
			}
		}
		return nil, &custerr.CustomError{
			Err:     err,
			ErrCode: http.StatusInternalServerError,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, &custerr.CustomError{
			Err:     fmt.Errorf("invalid username & password"),
			ErrCode: http.StatusBadRequest,
		}
	}

	return m.generateAuthToken(*user)
}

func (m Module) RefreshToken(ctx context.Context, req request.RefreshTokenRequest) (*response.AuthResponse, error) {
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &custerr.CustomError{
				Err:     fmt.Errorf("invalid refresh token"),
				ErrCode: http.StatusBadRequest,
			}
		}

		return []byte(m.authConfig.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uId := claims["uid"].(string)
		userId, err := primitive.ObjectIDFromHex(uId)
		if err != nil {
			return nil, &custerr.CustomError{
				Err:     fmt.Errorf("invalid refresh token"),
				ErrCode: http.StatusBadRequest,
			}
		}

		user, err := m.userRepo.FindById(ctx, userId)
		if err != nil {
			if err == dao.ErrNotFound {
				return nil, &custerr.CustomError{
					Err:     err,
					ErrCode: http.StatusBadRequest,
				}
			}
			return nil, &custerr.CustomError{
				Err:     err,
				ErrCode: http.StatusInternalServerError,
			}
		}

		return m.generateAuthToken(*user)
	}

	return nil, &custerr.CustomError{
		Err:     fmt.Errorf("invalid refresh token"),
		ErrCode: http.StatusBadRequest,
	}
}

func (m Module) generateAuthToken(user model.User) (*response.AuthResponse, error) {
	jwtData := model.JwtData{
		UserId:      user.ID.Hex(),
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Role:        user.Role,
		Permissions: user.Permissions,
	}

	expireAt := time.Now().Add(m.authConfig.AccessTokenDuration)
	stdClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireAt),
		ID:        user.ID.Hex(),
		Subject:   user.ID.Hex(),
	}

	jwtClaim := model.JwtClaim{
		RegisteredClaims: stdClaims,
		JwtData:          jwtData,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwtClaim)
	secret := []byte(m.authConfig.JwtSecret)

	accessToken, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	rToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := rToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = time.Now().Add(m.authConfig.RefreshTokenDuration).Unix()
	rtClaims["uid"] = user.ID.Hex()

	refreshToken, err := rToken.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}