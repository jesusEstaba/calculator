package domain

import "errors"

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive            = "inactive"
)

type User struct {
	ID       string
	Username string
	Password string
	Balance  float64
	Status   UserStatus
}

func (u *User) Withdraw(amount float64) error {
	if amount > u.Balance {
		return errors.New("insufficient balance")
	}
	u.Balance -= amount
	return nil
}
