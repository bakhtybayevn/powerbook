package dto

type RegisterUserRequest struct {
    Email       string `json:"email" example:"test@example.com"`
    DisplayName string `json:"display_name" example:"John"`
    Password    string `json:"password" example:"123456"`
}

type RegisterUserResponse struct {
    ID          string `json:"id"`
    Email       string `json:"email"`
    DisplayName string `json:"display_name"`
}
