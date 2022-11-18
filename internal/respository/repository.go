package respository

import (
	"avito_task/internal/entity"
	//"avito_task/internal/requests"
	"avito_task/pkg/dbcontext"
	"avito_task/pkg/log"
	"context"
)

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func (r repository) GetDeposit(ctx context.Context, ownerId int64) (entity.Deposit, error) {
	var deposit entity.Deposit
	err := r.db.With(ctx).Select().Model(ownerId, &deposit)
	return deposit, err
}

func (r repository) UpdateDeposit(ctx context.Context, deposit entity.Deposit) error {
	return r.db.With(ctx).Model(&deposit).Update()
}

func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) AddFunds(ctx context.Context, ownerId int64, amount int64) error {
	// обратиться к бд c запросом на добавление
	// средств по ownerId на amount
}

func (r repository) ReserveFunds(ctx context.Context, ownerId, serviceId, orderId, amount int64) error {
	// списать amount со счета ownerId
	// if err != nil return
	// создать transaction c ownerId, serviceId, orderId, amount
	// и IsPending = true
	// if err != nil вернуть amount на счет ownerId
}

func (r repository) ApplyReserve(ctx context.Context, ownerId, serviceId, orderId, amount int64) error {
	// найти transaction по ownerId, serviceId, orderId, amount
	// transaction.amount -= amount
	// if amount = 0 IsPending = false
	// данных добавление в отчет
}

func (r repository) GetBalance(ctx context.Context, ownerId int64) (balance int64, err error) {
	// обратиться к бд, вернуть результат
}

//func (r repository) GetHistory(ctx context.Context, req requests.GetHistoryRequest) ([]entity.Transaction, error) {
//	if err := req.Validate(); err != nil {
//		return nil, err
//	}
//
//	// if limit not specified, set equal to -1(meaning no limit in SQL)
//	if req.Limit == 0 {
//		req.Limit = -1
//	}
//
//	ownerUUID := uuid.MustParse(req.OwnerId)
//
//	return r.repo.GetForUser(ctx, ownerUUID, req.OrderBy, req.OrderDirection, req.Offset, req.Limit)
//}
//
