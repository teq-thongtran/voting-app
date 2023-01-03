package repository

import (
	"context"

	"gorm.io/gorm"

	"myapp/repository/poll"
	pollopt "myapp/repository/poll_option"
	"myapp/repository/user"
	uspo "myapp/repository/user_poll"
	usvo "myapp/repository/user_vote"
)

type Repository struct {
	User       user.Repository
	Poll       poll.Repository
	PollOption pollopt.Repository
	UserPoll   uspo.Repository
	UserVote   usvo.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User:       user.NewPG(getClient),
		Poll:       poll.NewPG(getClient),
		PollOption: pollopt.NewPG(getClient),
		UserPoll:   uspo.NewPG(getClient),
		UserVote:   usvo.NewPG(getClient),
	}
}
