package domain

type RandomStringRepository interface {
	Generate() (string, error)
}
