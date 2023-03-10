package model

import (
	"time"
)

type UserPoll struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	User      User      `json:"-"`
	PollID    int64     `json:"poll_id"`
	Poll      Poll      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
