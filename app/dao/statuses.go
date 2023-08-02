package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (s *status) SaveStatus(ctx context.Context, status *object.Status) error {
	_, err := s.db.ExecContext(ctx, "insert into status (id, account_id, content) values (?, ?, ?)",
		status.ID, status.AccountID, status.Content)
	if err != nil {
		return fmt.Errorf("failed to insert status into db: %w", err)
	}

	return nil
}
