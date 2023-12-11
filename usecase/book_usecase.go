package usecase

import (
	"context"
	"heji-server/domain"
)

type bookUseCase struct {
	bookUseCase domain.BookUseCase
}

func (b bookUseCase) Create(c context.Context, book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func NewBookUseCase(useCase domain.BookUseCase) domain.BookUseCase {
	return &bookUseCase{useCase}
}
