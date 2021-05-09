package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Language struct {
	ID          primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Name        string             `bson:"name, omitempty" json:"name,omitempty"`
	CountryCode string             `bson:"CountryCode, omitempty" json:"country_code,omitempty"`
}
