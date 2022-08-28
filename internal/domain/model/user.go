package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username    string             `bson:"username" json:"username"`
	FirstName   string             `bson:"first_name" json:"first_name"`
	LastName    string             `bson:"last_name" json:"last_name"`
	Password    string             `bson:"password" json:"-"`
	Role        UserRole           `bson:"role" json:"role"`
	Permissions []string           `bson:"permissions" json:"permissions"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	CreatedBy   string             `bson:"created_by" json:"created_by"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	UpdatedBy   string             `bson:"updated_by" json:"updated_by"`
	IsDeleted   bool               `bson:"is_deleted" json:"is_deleted"`
}

type UserUpdate struct {
	Username    string   `bson:"username,omitempty" json:"username"`
	FirstName   string   `bson:"first_name,omitempty" json:"first_name"`
	LastName    string   `bson:"last_name,omitempty" json:"last_name"`
	Password    string   `bson:"password,omitempty" json:"-"`
	Role        UserRole `bson:"role,omitempty" json:"role"`
	Permissions []string `bson:"permissions,omitempty" json:"permissions,omitempty"`
}
