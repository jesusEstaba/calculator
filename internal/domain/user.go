package domain

import "errors"

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Withdraw(amount float64) error {
	if amount > u.Balance {
		return errors.New("insufficient balance")
	}
	u.Balance -= amount
	return nil
}
