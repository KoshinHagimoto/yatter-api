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
		WHERE 
			(:max_id = -1 or s.id <= :max_id)
		AND 
			(:since_id = -1 or s.id >= :since_id)
		ORDER BY s.id DESC
		LIMIT :limit
	`

	property := map[string]interface{}{
		"max_id":   -1,
		"since_id": -1,
		"limit":    object.MaxLimit,
	}

	if timeline.MaxID != nil {
		property["max_id"] = *timeline.MaxID
	}

	if timeline.SinceID != nil {
		property["since_id"] = *timeline.SinceID
	}

	if timeline.Limit != nil {
		property["limit"] = *timeline.Limit
	}

	rows, err := t.db.NamedQueryContext(ctx, query, property)
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

func (t *timeline) GetHomeTimeline(ctx context.Context, accountID int64, timeline *object.Timeline) ([]*object.Status, error) {
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
			a.ID IN (
				SELECT
					following_id
				FROM 
					relationship
				WHERE
					follower_id = :account_id
			)
		AND
			(:max_id = -1 or s.id <= :max_id)
		AND 
			(:since_id = -1 or s.id >= :since_id)
		ORDER BY s.id DESC
		LIMIT :limit
	`

	property := map[string]interface{}{
		"account_id": accountID,
		"max_id":     -1,
		"since_id":   -1,
		"limit":      object.MaxLimit,
	}

	if timeline.MaxID != nil {
		property["max_id"] = *timeline.MaxID
	}

	if timeline.SinceID != nil {
		property["since_id"] = *timeline.SinceID
	}

	if timeline.Limit != nil {
		property["limit"] = *timeline.Limit
	}

	rows, err := t.db.NamedQueryContext(ctx, query, property)
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
