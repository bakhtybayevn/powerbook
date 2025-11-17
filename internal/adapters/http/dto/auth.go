package dto

type LoginRequest struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"123456"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
