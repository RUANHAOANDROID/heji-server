package repository

import (
	"context"
	"heji-server/domain"
	"heji-server/mongo"
)

type billRepository struct {
	database   mongo.Database
	collection string
}

func (b billRepository) Save(c context.Context, book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b billRepository) Delete(c context.Context, bid string) error {
	//TODO implement me
	panic("implement me")
}

func (b billRepository) List(c context.Context, userId string) (*[]domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b billRepository) Update(c context.Context, book *domain.Book) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewBillRepository(db mongo.Database, coll string) domain.BillRepository {
	return &billRepository{
		database:   db,
		collection: coll,
	}

}
