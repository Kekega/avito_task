package entity

// Deposit represents a user's account in the database.
type Deposit struct {
	OwnerId int64 `json:"owner_id" db:"pk"`
	Balance int64     `json:"balance"`
}
