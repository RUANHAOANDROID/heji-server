package usecase

import (
	"context"
	"heji-server/domain"
	"heji-server/internal/tokenutil"
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

func (uc *userUserCase) GetByTel(c context.Context, tel string) (domain.User, error) {
	user, err := uc.userRepository.GetByTel(c, tel)
	return user, err
}

func (uc *userUserCase) CreateAccessToken(user *domain.User, secret string, expiry time.Duration) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (uc *userUserCase) CreateRefreshToken(user *domain.User, secret string, expiry time.Duration) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
