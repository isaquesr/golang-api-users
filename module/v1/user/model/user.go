package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Model
type Users struct {
	// ID this is _id
	ID primitive.ObjectID `json:"_id" bson:"_id" swaggertype:"string"`

	// Name user
	Name string `json:"name" bson:"name" swaggertype:"string"`

	// Age user
	Age int `json:"age" bson:"age" swaggertype:"integer"`

	// Email user
	Email string `json:"email" bson:"email" swaggertype:"string"`

	// Password user
	Password string `json:"password" bson:"password" swaggertype:"string"`

	// Address user
	Address string `json:"address" bson:"address" swaggertype:"string"`
} //@name Response

// User CreateUser
type CreateUser struct {
	// Name user
	Name string `json:"name" valid:"-" bson:"name" swaggertype:"string"`

	// Age user
	Age int `json:"age" valid:"numeric" bson:"age" swaggertype:"integer"`

	// Email user
	Email string `json:"email" valid:"email" bson:"email" swaggertype:"string"`

	// Password user
	Password string `json:"password" valid:"type(string)" bson:"password" swaggertype:"string"`

	// Address user
	Address string `json:"address" valid:"type(string)" bson:"address" swaggertype:"string"`
}

// User UpdateUser
type UpdateUser struct {
	// Name user
	Name string `json:"name" bson:"name" swaggertype:"string"`

	// Age user
	Age int `json:"age" bson:"age" swaggertype:"integer"`

	// Email user
	Email string `json:"email" bson:"email" swaggertype:"string"`

	// Password user
	Password string `json:"password" bson:"password" swaggertype:"string"`

	// Address user
	Address string `json:"address" bson:"address" swaggertype:"string"`
}
