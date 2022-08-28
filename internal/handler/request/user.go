package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCreateRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type UserUpdateRequest struct {
	UserId    primitive.ObjectID `json:"user_id"`
	Username  string             `json:"username"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Password  string             `json:"password"`
	Role      string             `json:"role"`
}
