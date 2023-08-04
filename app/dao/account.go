package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

func (a *account) SaveAccount(ctx context.Context, account *object.Account) (accountID int64, err error) {
	result, err := a.db.ExecContext(ctx, "insert into account (username, password_hash) values (?, ?)",
		account.Username, account.PasswordHash)
	if err != nil {
		return 0, fmt.Errorf("failed to insert account into db: %w", err)
	}

	accountID, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	return accountID, nil
}

func (a *account) UpdateAccount(ctx context.Context, account *object.Account) error {
	query := `
		UPDATE account
		SET display_name = IFNULL(?, display_name),
			avatar = IFNULL(?, avatar),
			header = IFNULL(?, header),
			note = IFNULL(?, note)
		WHERE id = ?
	`

	_, err := a.db.ExecContext(ctx, query, account.DisplayName, account.Avatar, account.Header, account.Note, account.ID)
	if err != nil {
		return fmt.Errorf("failed to update account in db: %w", err)
	}

	return nil
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	// StructScanで取得したデータをentity（object.Account）にマッピング
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}
