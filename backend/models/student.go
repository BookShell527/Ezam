package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Student struct {
	Id       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Nrp      string        `bson:"nrp,omitempty" json:"nrp"`
	Name     string        `bson:"name,omitempty" json:"name"`
	Email    string        `bson:"email,omitempty" json:"email"`
	Password string        `bson:"password,omitempty" json:"-"`
}

type StudentRegisterDTO struct {
	Nrp      string `json:"nrp" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type StudentLoginDTO struct {
	Identifier string `json:"identifier,omitempty"`
	Password   string `json:"password" binding:"required"`
}
