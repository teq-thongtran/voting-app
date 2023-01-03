package usvo

import (
	"context"
	"fmt"

	"myapp/customError"
	"myapp/model"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	"myapp/repository/poll"
	pollopt "myapp/repository/poll_option"
	usvo "myapp/repository/user_vote"
)

type UserVoteUseCase interface {
	Create(ctx context.Context, req *payload.CreateUserVoteRequest) (*presenter.UserVoteResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest, pollId int64) (*presenter.ListUserVoteResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

type UseCase struct {
	UserVoteRepo   usvo.Repository
	PollRepo       poll.Repository
	PollOptionRepo pollopt.Repository
}

func New(repo *repository.Repository) UserVoteUseCase {
	return &UseCase{
		UserVoteRepo:   repo.UserVote,
		PollOptionRepo: repo.PollOption,
		PollRepo:       repo.Poll,
	}
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateUserVoteRequest,
) (*presenter.UserVoteResponseWrapper, error) {
	if err := u.validateCreate(ctx, req); err != nil {
		return nil, err
	}

	myUserVote := &model.UserVote{
		PollOptionId: req.PollOptionId,
		UserId:       ctx.Value("user_id").(int64),
	}

	err := u.UserVoteRepo.Create(ctx, myUserVote)
	if err != nil {
		return nil, customError.ErrModelCreate(err)
	}

	return &presenter.UserVoteResponseWrapper{UserVote: myUserVote}, nil
}

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteRequest) error {
	myUserVote, err := u.UserVoteRepo.GetByID(ctx, req.ID)
	if err != nil {
		return customError.ErrModelGet(err, "UserVote")
	}

	err = u.UserVoteRepo.Delete(ctx, myUserVote, false)
	if err != nil {
		return customError.ErrModelDelete(err)
	}

	if ctx.Value("user_id").(int64) != myUserVote.UserId {
		return customError.ErrGetByPolicty()
	}

	return nil
}

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListRequest,
	pollId int64,
) (*presenter.ListUserVoteResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s", req.OrderBy))
	}

	myPoll, err := u.PollRepo.GetByID(ctx, pollId)
	if err != nil {
		return nil, customError.ErrModelGet(err, "Poll")
	}

	err = u.validatePoll(ctx, myPoll, ctx.Value("user_id").(int64))

	if err != nil {
		return nil, err
	}

	myUserVotes, total, err := u.UserVoteRepo.GetList(ctx, req.Page, req.Limit, conditions, pollId, order)
	if err != nil {
		return nil, customError.ErrModelGet(err, "UserVote")
	}

	if req.Page == 0 {
		req.Page = 1
	}
	return &presenter.ListUserVoteResponseWrapper{
		Poll:      myPoll,
		UserVotes: myUserVotes,
		Meta: map[string]interface{}{
			"page":  req.Page,
			"limit": req.Limit,
			"total": total,
		},
	}, nil
}
