package entities

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
