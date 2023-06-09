package domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation struct {
	ID   *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type string              `json:"type" bson:"type"`
	Cost float64             `json:"cost" bson:"cost"`
}

func (o *Operation) GetID() *primitive.ObjectID {
	return o.ID
}

func (o *Operation) SetID(i *primitive.ObjectID) {
	o.ID = i
}

func (o *Operation) UpsertKey() *bson.M {
	if o.ID == nil {
		return nil
	}

	return &bson.M{"_id": o.ID}
}
