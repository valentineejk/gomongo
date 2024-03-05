package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Category string             `bson:"category,omitempty" json:"category,omitempty"`
	Stock    string             `bson:"stock,omitempty" json:"stock,omitempty"`
	Price    float32            `bson:"price,omitempty" json:"price,omitempty"`
}

type CreateProductRequest struct {
	Name     string  `bson:"name,omitempty" json:"name,omitempty"`
	Category string  `bson:"category,omitempty" json:"category,omitempty"`
	Stock    string  `bson:"stock,omitempty" json:"stock,omitempty"`
	Price    float32 `bson:"price,omitempty" json:"price,omitempty"`
}
