package repository

import (
	"context"
	"heji-server/domain"
	"heji-server/mongo"
)

// 账本存储库
type bookRepository struct {
	database   mongo.Database
	collection string
}

func (b bookRepository) List(c context.Context, userId string) (*[]domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookRepository) Update(c context.Context, book *domain.Book) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
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
