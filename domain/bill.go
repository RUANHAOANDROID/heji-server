package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bill struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Type   int64              `bson:"type"`
	Banner string             `bson:"banner"`
	Users  []User             `bson:"users"`
}
type BillUseCase interface {
	SaveBill(c context.Context, book *Book) error
	BillList(c context.Context, book *[]Book) error
	DeleteBill(c context.Context, bookId string) error
	UpdateBill(c context.Context, book *Book) error
}

type BillRepository interface {
	Save(c context.Context, book *Book) error
	Delete(c context.Context, bid string) error
	List(c context.Context, userId string) (*[]Book, error)
	Update(c context.Context, book *Book) (*Book, error)
}
