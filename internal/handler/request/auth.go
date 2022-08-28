package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
