package domain

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

type User struct {
	ID       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string              `json:"username" bson:"username"`
	Password []byte              `json:"-" bson:"password"`
	Balance  float64             `json:"balance" bson:"balance"`
	Status   UserStatus          `json:"status" bson:"status"`
}

func (u *User) Withdraw(amount float64) error {
	if amount > u.Balance {
		return errors.New("insufficient balance")
	}
	u.Balance -= amount
	return nil
}

func (u *User) GetID() *primitive.ObjectID {
	return u.ID
}

func (u *User) SetID(i *primitive.ObjectID) {
	u.ID = i
}

func (u *User) UpsertKey() *bson.M {
	if u.ID == nil {
		return nil
	}

	return &bson.M{"_id": u.ID}
}
