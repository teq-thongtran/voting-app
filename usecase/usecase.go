package usecase

import (
	"myapp/repository"
	"myapp/usecase/poll"
	pollopt "myapp/usecase/poll_option"
	"myapp/usecase/user"
	uspo "myapp/usecase/user_poll"
	usvo "myapp/usecase/user_vote"
)

type UseCase struct {
	User       user.UserUserCase
	Poll       poll.PollUseCase
	PollOption pollopt.PollOptionUseCase
	UserPoll   uspo.UserPollUseCase
	UserVote   usvo.UserVoteUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User:       user.New(repo),
		Poll:       poll.New(repo),
		PollOption: pollopt.New(repo),
		UserPoll:   uspo.New(repo),
		UserVote:   usvo.New(repo),
	}
}
