package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// TODO: Add Other APIs
	SaveAccount(ctx context.Context, account *object.Account) (accountID int64, err error)
	UpdateAccount(ctx context.Context, account *object.Account) error
}
