package usecase

import (
	"context"
	"heji-server/domain"
	"time"
)

type bookUseCase struct {
	repository domain.BookRepository
}

func (b bookUseCase) CreateBook(c context.Context, book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b bookUseCase) BookList(c context.Context, book *[]domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b bookUseCase) DeleteBook(c context.Context, bookId string) error {
	//TODO implement me
	panic("implement me")
}

func (b bookUseCase) JoinBook(c context.Context, code string) error {
	//TODO implement me
	panic("implement me")
}

func (b bookUseCase) UpdateBook(c context.Context, book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b bookUseCase) SharedBook(c context.Context, bookId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookUseCase) Create(c context.Context, book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func NewBookUseCase(repository domain.BookRepository, timeout time.Duration) domain.BookUseCase {
	return &bookUseCase{repository}
}
