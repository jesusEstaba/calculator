package domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID                *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OperationID       string              `json:"operation_id" bson:"operation_id"`
	UserID            string              `json:"user_id" bson:"user_id"`
	Amount            float64             `json:"amount" bson:"amount"`
	UserBalance       float64             `json:"user_balance" bson:"user_balance"`
	OperationResponse any                 `json:"operation_response" bson:"operation_response"`
	Date              string              `json:"created_at" bson:"created_at"`
}

type RecordSearch struct {
	UserID     string `json:"user_id"`
	SearchTerm string `json:"search_term"`
	Page       uint   `json:"page"`
	PerPage    uint   `json:"per_page"`
	Sort       string `json:"sort"`
}

func (r *Record) GetID() *primitive.ObjectID {
	return r.ID
}

func (r *Record) SetID(i *primitive.ObjectID) {
	r.ID = i
}

func (r *Record) UpsertKey() *bson.M {
	if r.ID == nil {
		return nil
	}

	return &bson.M{"_id": r.ID}
}
