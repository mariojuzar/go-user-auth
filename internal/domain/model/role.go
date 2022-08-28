package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRole string

const (
	BasicUser  UserRole = "basic_user"
	UserAdmin  UserRole = "user_admin"
	SuperAdmin UserRole = "superadmin"
)

var (
	UserRoles = map[string]UserRole{
		string(SuperAdmin): SuperAdmin,
		string(UserAdmin):  UserAdmin,
		string(BasicUser):  BasicUser,
	}
)

type RolePermission struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Role       UserRole           `bson:"role" json:"role"`
	Permission string             `bson:"permission" json:"permission"`
	AccessAPI  string             `bson:"access_api" json:"access_api"`
	Method     string             `bson:"method" json:"method"`
}
