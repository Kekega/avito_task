package respository

import (
	"avito_task/internal/entity"
	"context"
)

type Repository interface {
	// AddFunds adds amount to ownerId's deposit
	AddFunds(ctx context.Context, ownerId int64, amount int64) error
	// ReserveFunds reserves amount ad ownerId's deposit
	ReserveFunds(ctx context.Context, ownerId, serviceId, orderId, amount int64) error
	// ApplyReserve sets orderId pending status to false
	// and adds information to report
	ApplyReserve(ctx context.Context, ownerId, serviceId, orderId, amount int64) error
	// GetBalance returns current not reserved ownerId's balance
	GetBalance(ctx context.Context, ownerId int64) (balance int64, err error)
	// Get returns the Deposit with the specified owner's UUID.
	Get(ctx context.Context, ownerId int64) (entity.Deposit, error)
	// Create saves a new Deposit in the storage.
	Create(ctx context.Context, deposit entity.Deposit) error
	// Update updates the changes to the given Deposit to db.
	Update(ctx context.Context, deposit entity.Deposit) error

	//GetHistory(ctx context.Context, ownerId, limit, offset int64) ([]entity.Transaction, error)
	//MonthlyReport(ctx context.Context, year, month int64) ([]entity.Transaction, error)
}
