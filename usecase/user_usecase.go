package usecase

import (
	"context"
	"heji-server/domain"
	"time"
)

type userUserCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (uc *userUserCase) Register(c context.Context, user *domain.User) error {
	err := uc.userRepository.Register(c, user)
	return err
}

func NewLoginUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUserCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (l userUserCase) GetByTel(c context.Context, email string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (l userUserCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	//TODO implement me
	panic("implement me")
}

func (l userUserCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	//TODO implement me
	panic("implement me")
}
