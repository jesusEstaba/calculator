package domain

type TokenRepository interface {
	Generate(userID string) (string, error)
	Verify(token string) (string, error)
}
