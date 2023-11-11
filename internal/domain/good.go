package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Good struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
	Quantity    int64              `json:"quantity" bson:"quantity"`
}
