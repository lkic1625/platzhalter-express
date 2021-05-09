package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Word struct {
	ID           primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt, omitempty" json:"created_at"`
	updatedAt    time.Time          `bson:"updatedAt, omitempty" json:"updated_at"`
	Synonyms     []string           `bson:"synonyms, omitempty" json:"synonyms,omitempty"`
	Antonyms     []string           `bson:"antonyms, omitempty" json:"antonyms,omitempty"`
	PartOfSpeech string             `bson:"PartOfSpeech, omitempty" json:"part_of_speech,omitempty"`
}
