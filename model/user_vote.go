package model

import (
	"time"
)

type UserVote struct {
	ID           int64      `json:"id"`
	UserID       int64      `json:"user_id"`
	User         User       `json:"-"`
	PollOptionID int64      `json:"poll_option_id"`
	PollOption   PollOption `json:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
