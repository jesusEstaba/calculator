package domain

type PasswordRepository interface {
	Generate(string) ([]byte, error)
	Compare(string, []byte) error
}
