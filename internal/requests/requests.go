package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Request represents a JSON data of an API request.
type Request interface {
	// Validate validates the request's fields.
	Validate() error
}

// GetBalanceRequest represents a request to get balance of specific user.
type GetBalanceRequest struct {
	OwnerId string `json:"owner_id"`
}

// Validate validates the GetBalanceRequest fields.
func (r GetBalanceRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.OwnerId, validation.Required),
	)
}

// AddFundsRequest represents a request to add money to user's balance.
type AddFundsRequest struct {
	OwnerId int64 `json:"owner_id"`
	Amount  int64 `json:"amount"`
}

func (r AddFundsRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.OwnerId, validation.Required),
		validation.Field(&r.Amount, validation.Required),
	)
}

// ReserveRequest represents a request to transfer money from one user to another.
type ReserveRequest struct {
	SenderId  string `json:"sender_id"`
	ServiceId string `json:"service_id"`
	Amount    int64  `json:"amount"`
	OrderId   int64  `json:"order_id"`
}

// Validate validates the ReserveRequest fields.
func (r ReserveRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.SenderId, validation.Required),
		validation.Field(&r.ServiceId, validation.Required),
		validation.Field(&r.Amount, validation.Required, validation.Min(0).Exclusive()),
		validation.Field(&r.OrderId, validation.Required),
	)
}


type ApplyReserveRequest struct {
	SenderId  string `json:"sender_id"`
	ServiceId string `json:"service_id"`
	Amount    int64  `json:"amount"`
	OrderId   int64  `json:"order_id"`
}

func (r ApplyReserveRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.SenderId, validation.Required),
		validation.Field(&r.ServiceId, validation.Required),
		validation.Field(&r.Amount, validation.Required, validation.Min(0).Exclusive()),
		validation.Field(&r.OrderId, validation.Required),
	)
}


// GetHistoryRequest represents a request to get a list of all user's transactions: top-ups, withdrawals and transfers.
type GetHistoryRequest struct {
	OwnerId        string `json:"owner_id"`
	Offset         int    `json:"offset,omitempty"`
	Limit          int    `json:"limit,omitempty"`
	OrderBy        string `json:"order_by,omitempty"`
	OrderDirection string `json:"order_direction,omitempty"`
}

// Validate validates the GetHistoryRequest.
func (r GetHistoryRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.OwnerId, validation.Required),
		validation.Field(&r.Offset, validation.Min(0)),
		validation.Field(&r.Limit, validation.Min(1)),
		validation.Field(&r.OrderBy, validation.In("transaction_date", "amount")),
		validation.Field(&r.OrderDirection, validation.In("ASC", "DESC")),
	)
}
