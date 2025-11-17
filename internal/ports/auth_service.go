package ports

// AuthService проверяет валидность JWT и извлекает userID
type AuthService interface {
    ParseToken(token string) (string, error)
}
