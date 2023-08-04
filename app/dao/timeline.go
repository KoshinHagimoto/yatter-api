package dao

import (
	"context"
	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (t *timeline) GetPublicTimeline(ctx context.Context, timeline *object.Timeline) ([]*object.Status, error) {
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
	`

	args := []interface{}{}

	if timeline.MaxID != nil {
		query += " WHERE s.id <= ?"
		args = append(args, *timeline.MaxID)
	}

	if timeline.SinceID != nil {
		if len(args) > 0 {
			query += " AND s.id >= ?"
		} else {
			query += " WHERE s.id >= ?"
		}
		args = append(args, *timeline.SinceID)
	}

	query += " ORDER BY s.id DESC"

	if timeline.Limit != nil {
		query += " LIMIT ?"
		args = append(args, *timeline.Limit)
	}

	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []*object.Status
	for rows.Next() {
		var s object.Status
		var a object.Account
		if err := rows.Scan(
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
			return nil, err
		}

		s.Account = &a
		statuses = append(statuses, &s)
	}

	return statuses, nil
}
