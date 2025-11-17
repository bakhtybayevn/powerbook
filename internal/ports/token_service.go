package ports

type TokenService interface {
    GenerateToken(userID string) (string, error)
}
