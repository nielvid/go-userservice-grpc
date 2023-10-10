package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FirstName string             `bson:"first_name" json:"first_name,omitempty"`
	LastName  string             `bson:"last_name" json:"last_name,omitempty"`
	PhoneNumber  string            `bson:"phone_number" json:"phone_number,omitempty"`
	Email  string             `bson:"email" json:"email,omitempty"`
	Password string             `bson:"password" json:"password,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
