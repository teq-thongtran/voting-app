package model

import (
	"gorm.io/gorm"
	"time"
)

type Poll struct {
	ID           int64           `json:"id"`
	UserId       int64           `json:"user_id"`
	User         User            `json:"-"`
	PollPolicy   int8            `json:"poll_policy"`
	PollTitle    string          `json:"poll_title"`
	PollVoteType int8            `json:"poll_vote_type"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `json:"-"`
}
