package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Model
type Users struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Age         int                `json:"age" bson:"age"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password"`
	Address     string             `json:"address" bson:"address"`
	CreatedDate string             `json:"createdDate" bson:"createdDate"`
	LastUpdate  string             `json:"lastUpdate" bson:"lastUpdate"`
}

// User CreateUser
type CreateUser struct {
	Name        string `json:"name" bson:"name"`
	Age         int    `json:"age" bson:"age"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password" bson:"password"`
	Address     string `json:"address" bson:"address"`
	CreatedDate string `json:"createdDate" bson:"createdDate"`
	LastUpdate  string `json:"lastUpdate" bson:"lastUpdate"`
}

// User UpdateUser
type UpdateUser struct {
	Name       string `json:"name" bson:"name"`
	Age        int    `json:"age" bson:"age"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	Address    string `json:"address" bson:"address"`
	LastUpdate string `json:"lastUpdate" bson:"lastUpdate"`
}
