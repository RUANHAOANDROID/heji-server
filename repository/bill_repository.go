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

func (b billRepository) Save(c context.Context, bill *domain.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (b billRepository) Delete(c context.Context, bid string) error {
	//TODO implement me
	panic("implement me")
}

func (b billRepository) List(c context.Context, bookId string) (*[]domain.Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (b billRepository) Update(c context.Context, book *domain.Bill) error {
	//TODO implement me
	panic("implement me")
}

func NewBillRepository(db mongo.Database, coll string) domain.BillRepository {
	return &billRepository{
		database:   db,
		collection: coll,
	}

}
