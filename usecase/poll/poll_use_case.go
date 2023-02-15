package poll

import (
	"context"
	"errors"
	"fmt"
	"myapp/customError"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	"myapp/repository/poll"
	pollopt "myapp/repository/poll_option"
	"myapp/repository/user"
	usvo "myapp/repository/user_vote"

	"gorm.io/gorm"

	"myapp/model"
)

type PollUseCase interface {
	Create(ctx context.Context, req *payload.CreatePollRequest) (*presenter.PollResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdatePollRequest) (*presenter.PollResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.PollResponseWrapper, error)
	GetDetailByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.PollDetailResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest) (*presenter.ListPollResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

type UseCase struct {
	PollRepo       poll.Repository
	UserRepo       user.Repository
	PollOptionRepo pollopt.Repository
	UserVoteRepo   usvo.Repository
}

func New(repo *repository.Repository) PollUseCase {
	return &UseCase{
		PollRepo:       repo.Poll,
		UserRepo:       repo.User,
		PollOptionRepo: repo.PollOption,
		UserVoteRepo:   repo.UserVote,
	}
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreatePollRequest,
) (*presenter.PollResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	userID := ctx.Value("user_id").(int64)
	myUser, err := u.UserRepo.GetByID(ctx, userID)

	if err != nil {
		return nil, customError.ErrModelGet(err, "User")
	}

	myPoll := &model.Poll{
		PollPolicy:   req.PollPolicy,
		PollTitle:    req.PollTitle,
		PollVoteType: req.PollVoteType,
		UserID:       myUser.ID,
	}

	err = u.PollRepo.Create(ctx, myPoll)
	if err != nil {
		return nil, customError.ErrModelCreate(err)
	}

	return &presenter.PollResponseWrapper{Poll: myPoll}, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdatePollRequest,
) (*presenter.PollResponseWrapper, error) {
	myPoll, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.PollRepo.Update(ctx, myPoll)
	if err != nil {
		return nil, customError.ErrModelUpdate(err)
	}

	return &presenter.PollResponseWrapper{Poll: myPoll}, nil
}

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteRequest) error {
	myPoll, err := u.PollRepo.GetByID(ctx, req.ID)
	if err != nil {
		return customError.ErrModelGet(err, "Poll")
	}

	if myPoll.UserID != ctx.Value("user_id").(int64) {
		return customError.ErrGetByPolicty()
	}

	err = u.PollRepo.Delete(ctx, myPoll, false)
	if err != nil {
		return customError.ErrModelDelete(err)
	}

	return nil
}

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListRequest,
) (*presenter.ListPollResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s", req.OrderBy))
	}

	myPolls, total, err := u.PollRepo.GetList(ctx, req.Page, req.Limit, conditions, order)
	if err != nil {
		return nil, customError.ErrModelGet(err, "Poll")
	}

	if req.Page == 0 {
		req.Page = 1
	}
	return &presenter.ListPollResponseWrapper{
		Polls: myPolls,
		Meta: map[string]interface{}{
			"page":  req.Page,
			"limit": req.Limit,
			"total": total,
		},
	}, nil
}

func (u *UseCase) GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.PollResponseWrapper, error) {
	myPoll, err := u.PollRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrModelNotFound()
		}

		return nil, customError.ErrModelGet(err, "Poll")
	}
	err = u.validatePoll(ctx, myPoll, ctx.Value("user_id").(int64))
	if err != nil {
		return nil, err
	}
	return &presenter.PollResponseWrapper{Poll: myPoll}, nil
}

func (u *UseCase) GetDetailByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.PollDetailResponseWrapper, error) {
	myPoll, err := u.PollRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrModelNotFound()
		}

		return nil, customError.ErrModelGet(err, "Poll")
	}
	err = u.validatePoll(ctx, myPoll, ctx.Value("user_id").(int64))
	if err != nil {
		return nil, err
	}

	var conditions = map[string]interface{}{}
	conditions["poll_id"] = req.ID
	myPollOptions, _, _ := u.PollOptionRepo.GetList(ctx, 1, 100, conditions, nil)
	myUserVotes, _, _ := u.UserVoteRepo.GetList(ctx, 1, 100, nil, req.ID, nil)
	return &presenter.PollDetailResponseWrapper{
		Poll:        myPoll,
		PollOptions: myPollOptions,
		UserVotes:   myUserVotes,
	}, nil
}
