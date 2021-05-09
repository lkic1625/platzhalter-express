package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notebook struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Name      string             `bson:"name, omitempty" json:"name,omitempty"`
	Words     []string           `bson:"words, omitempty" json:"words,omitempty"`
	Language  Language           `bson:"language, omitempty" json:"language"`
	CreatedAt time.Time          `bson:"createdAt, omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updateAt, omitempty" json:"updated_at"`
}
