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

func (s *status) SaveStatus(ctx context.Context, status *object.Status) (statusID int64, err error) {
	result, err := s.db.ExecContext(ctx, "insert into status (account_id, content) values (?, ?)",
		status.Account.ID, status.Content)
	if err != nil {
		return 0, fmt.Errorf("failed to insert status into db: %w", err)
	}

	statusID, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	return statusID, nil
}

func (s *status) DeleteStatus(ctx context.Context, statusID int64) error {
	_, err := s.db.ExecContext(ctx, "delete from status where id = ?", statusID)
	if err != nil {
		return fmt.Errorf("failed to delete status from db: %w", err)
	}

	return nil
}

func (r *status) FindStatusByID(ctx context.Context, statusID int64) (*object.Status, error) {
	query := `
		SELECT 
			s.id AS status_id,
			s.content,
			s.create_at AS status_created_at,
			a.id AS account_id,
			a.username,
			a.display_name,
			a.create_at AS account_created_at,
			a.note,
			a.avatar,
			a.header
		FROM 
			status AS s
		JOIN
			account AS a
		ON
			s.account_id = a.id
		WHERE
			s.id = ?
	`

	row := r.db.QueryRowContext(ctx, query, statusID)

	var s object.Status
	var a object.Account
	if err := row.Scan(
		&s.ID,
		&s.Content,
		&s.CreateAt,
		&a.ID,
		&a.Username,
		&a.DisplayName,
		&a.CreateAt,
		&a.Note,
		&a.Avatar,
		&a.Header,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	s.Account = &a

	return &s, nil
}
