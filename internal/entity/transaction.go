package entity

import (
	"time"
)

type Transaction struct {
	OwnerId         int64     `json:"owner_id"`
	Id              int64     `json:"id" db:"pk"`
	ServiceId       int64     `json:"service_id"`
	OrderId         int64     `json:"order_id"`
	Amount          int64     `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
	IsPending       bool      `json:"is_pending"`
}
