package usecase

import (
	"context"
	"heji-server/domain"
	"time"
)

type billUseCase struct {
	repository domain.BillRepository
}

func (b billUseCase) SaveBill(c context.Context, bill *domain.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (b billUseCase) BillList(c context.Context, bill *[]domain.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (b billUseCase) DeleteBill(c context.Context, billId string) error {
	//TODO implement me
	panic("implement me")
}

func (b billUseCase) UpdateBill(c context.Context, bill *domain.Bill) error {
	//TODO implement me
	panic("implement me")
}

func NewBillUseCase(br domain.BillRepository, timeout time.Duration) domain.BillUseCase {
	return &billUseCase{br}
}
