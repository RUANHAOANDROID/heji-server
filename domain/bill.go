package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate go run github.com/wolfogre/gtag/cmd/gtag -types Bill -tags bson .
type Bill struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"json:"_id"`
	BookId     string             `bson:"book_id" json:"book_id"`
	Money      int64              `bson:"money" json:"money"`
	Category   string             `bson:"category" json:"category"`
	CreateUser string             `bson:"create_user" json:"create_user"`
	CreateTime string             `bson:"create_time"json:"create_time"`
	UpdateTime string             `bson:"update_time"json:"update_time"`
	Remark     string             `bson:"remark"json:"remark"`
	Images     []string           `bson:"images"json:"images"`
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
