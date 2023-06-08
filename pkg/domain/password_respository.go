package domain

type PasswordRepository interface {
	Generate(string) (string, error)
}
