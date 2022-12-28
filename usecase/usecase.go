package usecase

import (
	"myapp/repository"
	"myapp/usecase/poll"
	pollopt "myapp/usecase/poll_option"
	"myapp/usecase/user"
)

type UseCase struct {
	User       user.UserUserCase
	Poll       poll.PollUseCase
	PollOption pollopt.PollOptionUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User:       user.New(repo),
		Poll:       poll.New(repo),
		PollOption: pollopt.New(repo),
	}
}
