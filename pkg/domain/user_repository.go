package domain

type UserRepository interface {
	GetUser(string) (User, error)
	FindByUsername(string) (*User, error)
	Save(user User) (User, error)
	UpdateUser(User) error
}
