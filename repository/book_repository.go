package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"heji-server/domain"
)

// 账本存储库
type bookRepository struct {
	database   mongo.Database
	collection string
}

func (b bookRepository) Create(c context.Context, book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func NewBookRepository(db mongo.Database, coll string) domain.BookRepository {
	return &bookRepository{
		database:   db,
		collection: coll,
	}
}
