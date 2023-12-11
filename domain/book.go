package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Type   int64              `bson:"type"`
	Banner string             `bson:"banner"`
	Users  []User             `bson:"users"`
}
type BookUseCase interface {
	//Create 创建账本
	Create(c context.Context, book *Book) error
}

type BookRepository interface {
	Create(c context.Context, book *Book) error
}
