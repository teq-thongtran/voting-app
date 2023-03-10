package model

import (
	"time"
)

type PollOption struct {
	ID         int64      `json:"id"`
	UserID     int64      `json:"user_id"`
	User       User       `json:"-"`
	PollID     int64      `json:"poll_id"`
	Poll       Poll       `json:"-"`
	OptionText string     `json:"option_text"`
	UserVotes  []UserVote `json:"-"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
