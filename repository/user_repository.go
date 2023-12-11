package repository

import (
	"context"
	"heji-server/domain"
	"heji-server/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func (u userRepository) Create(c context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Fetch(c context.Context) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
