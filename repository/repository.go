package repository

import (
	"context"
	"myapp/repository/poll"

	"gorm.io/gorm"

	"myapp/repository/user"
)

type Repository struct {
	User user.Repository
	Poll poll.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User: user.NewPG(getClient),
		Poll: poll.NewPG(getClient),
	}
}
