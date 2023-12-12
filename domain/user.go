package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserUseCase interface {
	Register(c context.Context, user *User) error
	GetByTel(c context.Context, tel string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}

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
	Register(c context.Context, user *User) error
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByID(c context.Context, id string) (User, error)
}
