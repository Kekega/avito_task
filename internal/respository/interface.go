package respository

import (
	"context"
)

type Repository interface {
	AddFunds(ctx context.Context, ownerId int64, amount int64) error
	ReserveFunds(ctx context.Context, ownerId, serviceId, orderId, amount int64) error
	ApplyReserve(ctx context.Context, ownerId, serviceId, orderId, amount int64) error
	GetBalance(ctx context.Context, ownerId int64) (balance int64, err error)
	//GetHistory(ctx context.Context, ownerId, limit, offset int64) ([]entity.Transaction, error)
	//MonthlyReport(ctx context.Context, year, month int64) ([]entity.Transaction, error)
}
