package usecase

import (
	"context"
	"heji-server/domain"
	"time"
)

type billUseCase struct {
	repository domain.BillRepository
}

func (b billUseCase) SaveBill(c context.Context, book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b billUseCase) BillList(c context.Context, book *[]domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b billUseCase) DeleteBill(c context.Context, bookId string) error {
	//TODO implement me
	panic("implement me")
}

func (b billUseCase) UpdateBill(c context.Context, book *domain.Book) error {
	//TODO implement me
	panic("implement me")
}

func NewBillUseCase(br domain.BillRepository, timeout time.Duration) domain.BillUseCase {
	return &billUseCase{br}
}
