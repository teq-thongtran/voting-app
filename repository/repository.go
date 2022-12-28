package repository

import (
	"context"

	"gorm.io/gorm"

	"myapp/repository/poll"
	"myapp/repository/poll_option"
	"myapp/repository/user"
)

type Repository struct {
	User       user.Repository
	Poll       poll.Repository
	PollOption poll_option.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User:       user.NewPG(getClient),
		Poll:       poll.NewPG(getClient),
		PollOption: poll_option.NewPG(getClient),
	}
}
