package model

import (
	"time"

	"gorm.io/gorm"
)

type Poll struct {
	ID           int64           `json:"id"`
	UserId       int64           `json:"user_id"`
	User         User            `json:"-"`
	PollPolicy   string          `json:"poll_policy"`
	PollTitle    string          `json:"poll_title"`
	PollVoteType string          `json:"poll_vote_type"`
	UserPolls    []UserPoll      `json:"-"`
	PollOptions  []PollOption    `json:"-"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `json:"-"`
}
