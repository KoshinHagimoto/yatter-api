package object

import "time"

type Status struct {
	ID int64 `json:"id,omitempty" db:"id"`

	AccountID int64 `json:"-" db:"account_id"`

	Content string `json:"content,omitempty"`

	Url *string `json:"url,omitempty"`

	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`

	Account *Account `json:"account,omitempty" db:"account"`
}
