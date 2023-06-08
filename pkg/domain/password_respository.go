package domain

type PasswordRepository interface {
	Generate(string) ([]byte, error)
}
