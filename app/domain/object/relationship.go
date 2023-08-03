package object

import "time"

type Relationship struct {
	ID          int64     `db:"id"`
	FollowerID  int64     `db:"follower_id"`
	FollowingID int64     `db:"following_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
