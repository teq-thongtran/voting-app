package usecase

import (
	"myapp/repository"
	"myapp/usecase/poll"
	"myapp/usecase/user"
)

type UseCase struct {
	User user.UserUserCase
	Poll poll.PollUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User: user.New(repo),
		Poll: poll.New(repo),
	}
}
