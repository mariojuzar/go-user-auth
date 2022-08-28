package response

type UserResponse struct {
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}
