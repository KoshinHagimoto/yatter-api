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
	var statuses []*object.Status
	query := "SELECT * FROM status"
	args := []interface{}{}

	if timeline.MaxID != nil {
		query += " WHERE id <= ?"
		args = append(args, *timeline.MaxID)
	}

	if timeline.SinceID != nil {
		if len(args) > 0 {
			query += " AND id >= ?"
		} else {
			query += " WHERE id >= ?"
		}
		args = append(args, *timeline.SinceID)
	}

	query += " ORDER BY id DESC"

	if timeline.Limit != nil {
		query += " LIMIT ?"
		args = append(args, *timeline.Limit)
	}

	if err := t.db.SelectContext(ctx, &statuses, query, args...); err != nil {
		return nil, err
	}

	return statuses, nil
}
