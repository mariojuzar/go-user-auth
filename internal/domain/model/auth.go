package model

import (
	"github.com/golang-jwt/jwt/v4"
)

const (
	JwtKey = "JWT_DATA_KEY"
)

type JwtData struct {
	UserId      string   `json:"user_id"`
	Username    string   `json:"username"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Role        UserRole `json:"role"`
	Permissions []string `json:"permissions"`
}

type JwtClaim struct {
	jwt.RegisteredClaims
	JwtData
}
