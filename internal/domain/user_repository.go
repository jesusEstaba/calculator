package domain

type UserRepository interface {
	GetUser(string) (User, error)
	UpdateUser(User) error
}
