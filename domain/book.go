package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollBook = "books" //mongo collection users
)

type Book struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Type      int64              `bson:"type" json:"type"`
	Banner    string             `bson:"banner" json:"banner"`
	CrtUserId string             `bson:"crt_user_id" json:"crt_user_id"`
	Users     []string           `bson:"users" json:"users"`
	IsInitial bool               `bson:"is_initial" json:"is_initial"`
}
type BookUseCase interface {
	CreateBook(c context.Context, book *Book) error
	BookList(c context.Context, userId string) (*[]Book, error)
	DeleteBook(c context.Context, bookId string) error
	JoinBook(c context.Context, code string) error
	UpdateBook(c context.Context, book *Book) error
	SharedBook(c context.Context, bookId string) (string, error)
}

type BookRepository interface {
	CreateOne(c context.Context, book *Book) error
	FindOne(c context.Context, id primitive.ObjectID) (Book, error)
	FindInitialBook(c context.Context, tel string) (Book, error)
	List(c context.Context, userId string) (*[]Book, error)
	Update(c context.Context, book *Book) (*Book, error)
}
