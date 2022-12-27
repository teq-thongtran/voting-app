package poll

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"myapp/customError"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	"myapp/repository/poll"
	"myapp/repository/user"
	"strings"

	"myapp/model"
)

type PollUseCase interface {
	Create(ctx context.Context, req *payload.CreatePollRequest) (*presenter.PollResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdatePollRequest) (*presenter.PollResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.PollResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest) (*presenter.ListPollResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

type UseCase struct {
	PollRepo poll.Repository
	UserRepo user.Repository
}

func New(repo *repository.Repository) PollUseCase {
	return &UseCase{
		PollRepo: repo.Poll,
		UserRepo: repo.User,
	}
}

func (u *UseCase) validateCreate(req *payload.CreatePollRequest) error {
	if req.PollTitle == "" {
		return customError.ErrRequestInvalidParam("poll_title")
	}

	req.PollTitle = strings.TrimSpace(req.PollTitle)
	if len(req.PollTitle) == 0 {
		req.PollTitle = ""
		return customError.ErrRequestInvalidParam("name_poll")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreatePollRequest,
) (*presenter.PollResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	userId := ctx.Value("user_id").(int64)
	myUser, err := u.UserRepo.GetByID(ctx, userId)

	if err != nil {
		return nil, customError.ErrModelGet(err, "User")
	}

	myPoll := &model.Poll{
		PollPolicy:   req.PollPolicy,
		PollTitle:    req.PollTitle,
		PollVoteType: req.PollVoteType,
		UserId:       myUser.ID,
	}

	err = u.PollRepo.Create(ctx, myPoll)
	if err != nil {
		return nil, customError.ErrModelCreate(err)
	}

	return &presenter.PollResponseWrapper{Poll: myPoll}, nil
}

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdatePollRequest) (*model.Poll, error) {
	myPoll, err := u.PollRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, customError.ErrModelGet(err, "Poll")
	}

	if req.PollTitle != nil {
		*req.PollTitle = strings.TrimSpace(*req.PollTitle)
		if len(*req.PollTitle) == 0 {
			return nil, customError.ErrRequestInvalidParam("name")
		}

		myPoll.PollTitle = *req.PollTitle
	}

	userId := ctx.Value("user_id").(int64)

	if myPoll.UserId != userId {
		return nil, customError.ErrUnauthorized(nil)
	}

	return myPoll, nil
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
	conditions["user_id"] = ctx.Value("user_id")
	myPolls, total, err := u.PollRepo.GetList(ctx, req.Search, req.Page, req.Limit, conditions, order)
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

	return &presenter.PollResponseWrapper{Poll: myPoll}, nil
}
