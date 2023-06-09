package entities

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
