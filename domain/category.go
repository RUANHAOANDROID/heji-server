package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID     primitive.ObjectID `bson:"_id"`
	BookID string             `bson:"book_id"`
	Type   int                `bson:"type"`
	Level  int                `bson:"level"`
}
