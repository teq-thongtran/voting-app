package uspo

import (
	"context"
	"fmt"

	"myapp/customError"
	"myapp/model"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	"myapp/repository/poll"
	"myapp/repository/user"
	uspo "myapp/repository/user_poll"
)

type UserPollUseCase interface {
	Create(ctx context.Context, req *payload.CreateUserPollRequest) (*presenter.UserPollResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest, pollID int64) (*presenter.ListUserPollResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

type UseCase struct {
	UserPollRepo uspo.Repository
	UserRepo     user.Repository
	PollRepo     poll.Repository
}

func New(repo *repository.Repository) UserPollUseCase {
	return &UseCase{
		UserPollRepo: repo.UserPoll,
		UserRepo:     repo.User,
		PollRepo:     repo.Poll,
	}
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateUserPollRequest,
) (*presenter.UserPollResponseWrapper, error) {
	if err := u.validateCreate(ctx, req); err != nil {
		return nil, err
	}

	myUserPoll := &model.UserPoll{
		PollID: req.PollID,
		UserID: req.UserID,
	}

	err := u.UserPollRepo.Create(ctx, myUserPoll)
	if err != nil {
		return nil, customError.ErrModelCreate(err)
	}

	return &presenter.UserPollResponseWrapper{UserPoll: myUserPoll}, nil
}

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteRequest) error {
	myUserPoll, err := u.UserPollRepo.GetByID(ctx, req.ID)
	if err != nil {
		return customError.ErrModelGet(err, "UserPoll")
	}

	myPoll, err := u.PollRepo.GetByID(ctx, myUserPoll.PollID)
	if err != nil {
		return customError.ErrModelGet(err, "Poll")
	}

	err = u.validatePoll(myPoll, ctx.Value("user_id").(int64))
	if err != nil {
		return err
	}

	err = u.UserPollRepo.Delete(ctx, myUserPoll, false)
	if err != nil {
		return customError.ErrModelDelete(err)
	}

	return nil
}

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListRequest,
	pollID int64,
) (*presenter.ListUserPollResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s", req.OrderBy))
	}

	myPoll, err := u.PollRepo.GetByID(ctx, pollID)
	if err != nil {
		return nil, customError.ErrModelGet(err, "Poll")
	}

	err = u.validatePoll(myPoll, ctx.Value("user_id").(int64))

	if err != nil {
		return nil, err
	}

	conditions["poll_id"] = pollID
	myUserPolls, total, err := u.UserPollRepo.GetList(ctx, req.Page, req.Limit, conditions, order)
	if err != nil {
		return nil, customError.ErrModelGet(err, "UserPoll")
	}

	if req.Page == 0 {
		req.Page = 1
	}
	return &presenter.ListUserPollResponseWrapper{
		Poll:      myPoll,
		UserPolls: myUserPolls,
		Meta: map[string]interface{}{
			"page":  req.Page,
			"limit": req.Limit,
			"total": total,
		},
	}, nil
}
