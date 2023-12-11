package usecase

import (
	"context"
	"heji-server/domain"
	"time"
)

type loginUserCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (l loginUserCase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (l loginUserCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	//TODO implement me
	panic("implement me")
}

func (l loginUserCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	//TODO implement me
	panic("implement me")
}

func NewLoginUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUseCase {
	return &loginUserCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
