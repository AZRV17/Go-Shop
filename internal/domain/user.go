package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Login    string             `json:"login" bson:"login"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}
