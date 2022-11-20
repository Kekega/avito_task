package transaction

import (
	"context"
	"strconv"
	"time"

	"avito_task/internal/entity"
	"avito_task/internal/requests"
	"avito_task/pkg/log"
)

// Service encapsulates usecase logic for transactions.
type Service interface {
	// CreateUpdateTransaction creates a Transaction based on AddFundsRequest.
	CreateUpdateTransaction(ctx context.Context, req requests.AddFundsRequest) (Transaction, error)
	// CreateTransferTransaction creates a Transaction based on ReserveRequest.
	CreateTransferTransaction(ctx context.Context, req requests.ReserveRequest) (Transaction, error)
	// GetHistory returns a list of all transactions related to the user with the given ID.
	GetHistory(ctx context.Context, req requests.GetHistoryRequest) ([]entity.Transaction, error)
	// Count returns a number of all Transactions in the database. Mainly used for testing purposes.
	Count(ctx context.Context) (int64, error)
}

// Transaction represents the data about a transaction.
type Transaction struct {
	entity.Transaction
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new Transaction service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) CreateUpdateTransaction(ctx context.Context, req requests.AddFundsRequest) (Transaction, error) {
	if err := req.Validate(); err != nil {
		return Transaction{}, err
	}

	ownerId := req.OwnerId
	tx := entity.Transaction{
		//ServiceId:       req.OwnerId,
		TransactionDate: time.Now().UTC(),
	}
	tx.OwnerId = ownerId
	if req.Amount < 0 {
		tx.Amount = -req.Amount
	} else {
		tx.Amount = req.Amount
	}

	err := s.repo.Create(ctx, &tx)
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{tx}, err
}

func (s service) CreateTransferTransaction(ctx context.Context, req requests.ReserveRequest) (Transaction, error) {
	if err := req.Validate(); err != nil {
		return Transaction{}, err
	}

	tx := entity.Transaction{
		Id:              0, // auto-incremented
		OwnerId:        req.SenderId,
		ServiceId:     req.ServiceId,
		Amount:          req.Amount,
		OrderId:     req.OrderId,
		TransactionDate: time.Now().UTC(),
	}

	err := s.repo.Create(ctx, &tx)
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{tx}, err
}

func (s service) GetHistory(ctx context.Context, req requests.GetHistoryRequest) ([]entity.Transaction, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// if limit not specified, set equal to -1(meaning no limit in SQL)
	if req.Limit == 0 {
		req.Limit = -1
	}

	ownerUUID, err := strconv.ParseInt(req.OwnerId, 10, 64)
	if err != nil {
		return nil, err
	}

	return s.repo.GetForUser(ctx, ownerUUID, req.OrderBy, req.OrderDirection, req.Offset, req.Limit)
}

func (s service) Count(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}
