package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// LoginRequest 登录请求
type LoginRequest struct {
	Tel      string `json:"tel" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录返回
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// UserUseCase 用户的用例
type UserUseCase interface {
	Register(c context.Context, user *User) error
	GetByTel(c context.Context, tel string) (User, error)
	CreateAccessToken(user *User, secret string, expiry time.Duration) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry time.Duration) (refreshToken string, err error)
}

const (
	CollUser = "users" //mongo collection users
)

// User Mongo 用户结构体
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name" json:"name"`
	Tel      string             `bson:"tel" json:"tel"`
	Password string             `bson:"password" json:"password"`
	ImageUrl string             `bson:"image_url"`
}

// UserRepository 定义用户资料接口
type UserRepository interface {
	Register(c context.Context, user *User) error
	GetByTel(c context.Context, tel string) (User, error)
}
