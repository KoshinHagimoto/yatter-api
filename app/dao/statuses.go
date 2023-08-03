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
	status struct {
		db *sqlx.DB
	}
)

// Repository for statuses　インターフェース満たすように生成
// daoでもrepositoryで定義されたメソッドを実装しなければならない
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (s *status) SaveStatus(ctx context.Context, status *object.Status) error {
	_, err := s.db.ExecContext(ctx, "insert into status (account_id, content) values (?, ?)",
		status.AccountID, status.Content)
	if err != nil {
		return fmt.Errorf("failed to insert status into db: %w", err)
	}

	return nil
}

func (s *status) DeleteStatus(ctx context.Context, statusID int64) error {
	_, err := s.db.ExecContext(ctx, "delete from status where id = ?", statusID)
	if err != nil {
		return fmt.Errorf("failed to delete status from db: %w", err)
	}

	return nil
}

func (r *status) FindStatusByID(ctx context.Context, statusID int64) (*object.Status, error) {
	entity := new(object.Status)
	// StructScanで取得したデータをentity（object.Status）にマッピング
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", statusID).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}
