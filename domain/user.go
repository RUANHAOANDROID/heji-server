package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollUser = "users"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Tel      string             `bson:"tel"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	ImageUrl string             `bson:"image_url"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByID(c context.Context, id string) (User, error)
}
type UserUseCase interface {
	Create(c context.Context, user *User) error
}
