package person

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person is potential Soylent Green
type Person struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Firstname string             `bson:"firstname,omitempty"`
	Lastname  string             `bson:"lastname,omitempty"`
}
