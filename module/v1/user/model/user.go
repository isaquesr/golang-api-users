package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Model
type Users struct {
	// ID this is _id
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id" swaggertype:"string"`

	// Name user
	Name string `json:"name,omitempty" bson:"name" swaggertype:"string"`

	// Age user
	Age int `json:"age,omitempty" bson:"age" swaggertype:"integer"`

	// Email user
	Email string `json:"email,omitempty" bson:"email" swaggertype:"string"`

	// Password user
	Password string `json:"password,omitempty" bson:"password" swaggertype:"string"`

	// Address user
	Address string `json:"address,omitempty" bson:"address" swaggertype:"string"`

	// CreatedDate user
	CreatedDate string `json:"createdDate,omitempty" bson:"createdDate" swaggertype:"string"`

	// LastUpdate user
	LastUpdate string `json:"lastUpdate,omitempty" bson:"lastUpdate" swaggertype:"string"`
} //@name Response

// User CreateUser
type CreateUser struct {
	// Name user
	Name string `json:"name,omitempty" bson:"name" swaggertype:"string"`

	// Age user
	Age int `json:"age,omitempty" bson:"age" swaggertype:"integer"`

	// Email user
	Email string `json:"email,omitempty" bson:"email" swaggertype:"string"`

	// Password user
	Password string `json:"password,omitempty" bson:"password" swaggertype:"string"`

	// Address user
	Address string `json:"address,omitempty" bson:"address" swaggertype:"string"`

	// CreatedDate user
	CreatedDate string `json:"createdDate,omitempty" bson:"createdDate" swaggertype:"string"`

	// LastUpdate user
	LastUpdate string `json:"lastUpdate,omitempty" bson:"lastUpdate" swaggertype:"string"`
}

// User UpdateUser
type UpdateUser struct {
	// Name user
	Name string `json:"name,omitempty" bson:"name" swaggertype:"string"`

	// Age user
	Age int `json:"age,omitempty" bson:"age" swaggertype:"integer"`

	// Email user
	Email string `json:"email,omitempty" bson:"email" swaggertype:"string"`

	// Password user
	Password string `json:"password,omitempty" bson:"password" swaggertype:"string"`

	// Address user
	Address string `json:"address,omitempty" bson:"address" swaggertype:"string"`

	// LastUpdate user
	LastUpdate string `json:"lastUpdate,omitempty" bson:"lastUpdate" swaggertype:"string"`
}
