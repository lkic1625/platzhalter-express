package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name, omitempty" json:"name,omitempty"`
	Email    string             `bson:"email, omitempty" json:"email,omitempty"`
	Password string             `bson:"password, omitempty" json:"password,omitempty"`
}
