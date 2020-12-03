package scout

import "go.mongodb.org/mongo-driver/bson/primitive"

type Scout struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Firstname string             `bson:"firstname,omitempty" json:"firstname,omitempty"`
	Rank      string             `bson:"rank,omitempty" json:"rank,omitempty"`
	Den       int                `bson:"den,omitempty" json:"den,omitempty"`
	Address   string             `bson:"address,omitempty" json:"address,omitempty"`
	Parents   string             `bson:"parents,omitempty" json:"parents,omitempty"`
	Phone     string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
}
