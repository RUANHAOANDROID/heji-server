package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollBill = "bills" //mongo collection users
)

//go:generate go run github.com/wolfogre/gtag/cmd/gtag -types Bill -tags bson .
type Bill struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	BookId     primitive.ObjectID `bson:"book_id" json:"book_id"`
	Money      string             `bson:"money" json:"money"`
	Type       int                `bson:"type" json:"type"`
	Category   string             `bson:"category" json:"category"`
	CreateUser string             `bson:"create_user" json:"create_user"`
	CreateTime int64              `bson:"create_time" json:"create_time"`
	UpdateTime int64              `bson:"update_time" json:"update_time"`
	Remark     string             `bson:"remark" json:"remark"`
	Images     []string           `bson:"images" json:"images"`
}
type BillUseCase interface {
	SaveBill(c context.Context, bill *Bill) error
	BillList(c context.Context, bookId string, pageNumber, pageSize int64, bill *[]Bill) error
	DeleteBill(c context.Context, billId string) error
	UpdateBill(c context.Context, bill *Bill) error
}

type BillRepository interface {
	Save(c context.Context, bill *Bill) error
	Delete(c context.Context, bid string) error
	List(c context.Context, bookId string, pageNumber, pageSize int64, bills *[]Bill) error
	Update(c context.Context, book *Bill) error
}
