package deposit

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/google/uuid"
	"avito_task/internal/entity"
	"avito_task/internal/errors"
	"avito_task/internal/requests"
	"avito_task/pkg/log"
)

// Service encapsulates usecase logic for deposits.
type Service interface {
	GetBalance(ctx context.Context, req requests.GetBalanceRequest) (float32, error)
	Update(ctx context.Context, req requests.AddFundsRequest) error
	Transfer(ctx context.Context, req requests.ReserveRequest) error
	Count(ctx context.Context) (int64, error)
}

// Deposit represents the data about a deposit.
type Deposit struct {
	entity.Deposit
}

// Transaction represents the data about a transaction.
type Transaction struct {
	entity.Transaction
}

type service struct {
	repo            Repository
	logger          log.Logger
}

// NewService creates a new Deposit depositService.
func NewService(depositRepo Repository, logger log.Logger) Service {
	return service{depositRepo, logger}
}

func (s service) modifyBalance(ctx context.Context, ownerId int64, amount int64) error {
	dep, err := s.repo.Get(ctx, ownerId)

	// If deposit is not in DB yet, create it.
	if err == sql.ErrNoRows {
		dep = entity.Deposit{OwnerId: ownerId}
		if err = s.repo.Create(ctx, dep); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	dep.Balance += amount
	if dep.Balance < 0 {
		return errors.Forbidden("Insufficient funds to perform operation.")
	}

	return s.repo.Update(ctx, dep)
}

// GetBalance returns the balance of the Deposit whose owner whose OwnerId is equal to GetBalanceRequest.OwnerId.
func (s service) GetBalance(ctx context.Context, req requests.GetBalanceRequest) (float32, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	ownerId, err := strconv.ParseInt(req.OwnerId, 10, 64)
	if err != nil {
		return 0, err
	}

	deposit, err := s.repo.Get(ctx, ownerId)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	balance := float32(deposit.Balance)

	return balance, nil
}

// Update changes the balance of Deposit according to AddFundsRequest.
// It returns the Transaction which reflects the corresponding balance change in case of success.
func (s service) Update(ctx context.Context, req requests.AddFundsRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	ownerId := req.OwnerId
	if err := s.modifyBalance(ctx, ownerId, req.Amount); err != nil {
		return err
	}

	return nil
}

// Transfer sends money from one user to another according to ReserveRequest.
// It returns a Transaction which reflects the corresponding money transfer in case of success.
func (s service) Transfer(ctx context.Context, req requests.ReserveRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	senderUUID, recipientUUID := uuid.MustParse(req.SenderId), uuid.MustParse(req.ServiceId)
	if err := s.modifyBalance(ctx, senderUUID, -req.Amount); err != nil {
		return err
	}
	if err := s.modifyBalance(ctx, recipientUUID, req.Amount); err != nil {
		return err
	}

	return nil
}

// Count returns a number of Deposits in the database.
// Mainly used for testing purposes.
func (s service) Count(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}
