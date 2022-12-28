package pollopt

import (
	"context"
	"errors"
	"fmt"
	"myapp/repository/poll"

	"gorm.io/gorm"

	"myapp/customError"
	"myapp/model"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	pollopt "myapp/repository/poll_option"
	"myapp/repository/user"
)

type PollOptionUseCase interface {
	Create(ctx context.Context, req *payload.CreatePollOptionRequest) (*presenter.PollOptionResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdatePollOptionRequest) (*presenter.PollOptionResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.PollOptionResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest, pollId int64) (*presenter.ListPollOptionResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

type UseCase struct {
	PollOptionRepo pollopt.Repository
	UserRepo       user.Repository
	PollRepo       poll.Repository
}

func New(repo *repository.Repository) PollOptionUseCase {
	return &UseCase{
		PollOptionRepo: repo.PollOption,
		UserRepo:       repo.User,
		PollRepo:       repo.Poll,
	}
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreatePollOptionRequest,
) (*presenter.PollOptionResponseWrapper, error) {
	if err := u.validateCreate(ctx, req); err != nil {
		return nil, err
	}

	userId := ctx.Value("user_id").(int64)
	myUser, err := u.UserRepo.GetByID(ctx, userId)

	if err != nil {
		return nil, customError.ErrModelGet(err, "User")
	}

	myPollOption := &model.PollOption{
		OptionText: req.OptionText,
		PollId:     req.PollId,
		UserId:     myUser.ID,
	}

	err = u.PollOptionRepo.Create(ctx, myPollOption)
	if err != nil {
		return nil, customError.ErrModelCreate(err)
	}

	return &presenter.PollOptionResponseWrapper{PollOption: myPollOption}, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdatePollOptionRequest,
) (*presenter.PollOptionResponseWrapper, error) {
	myPollOption, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.PollOptionRepo.Update(ctx, myPollOption)
	if err != nil {
		return nil, customError.ErrModelUpdate(err)
	}

	return &presenter.PollOptionResponseWrapper{PollOption: myPollOption}, nil
}

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteRequest) error {
	myPollOption, err := u.PollOptionRepo.GetByID(ctx, req.ID)
	if err != nil {
		return customError.ErrModelGet(err, "PollOption")
	}

	err = u.validatePollOption(myPollOption, ctx.Value("user_id").(int64))
	if err != nil {
		return err
	}

	err = u.PollOptionRepo.Delete(ctx, myPollOption, false)
	if err != nil {
		return customError.ErrModelDelete(err)
	}

	return nil
}

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListRequest,
	pollId int64,
) (*presenter.ListPollOptionResponseWrapper, error) {
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

	err = u.validatePoll(myPoll, ctx.Value("user_id").(int64))

	if err != nil {
		return nil, err
	}

	conditions["poll_id"] = pollId
	myPollOptions, total, err := u.PollOptionRepo.GetList(ctx, req.Page, req.Limit, conditions, order)
	if err != nil {
		return nil, customError.ErrModelGet(err, "PollOption")
	}

	if req.Page == 0 {
		req.Page = 1
	}
	return &presenter.ListPollOptionResponseWrapper{
		PollOptions: myPollOptions,
		Meta: map[string]interface{}{
			"page":  req.Page,
			"limit": req.Limit,
			"total": total,
		},
	}, nil
}

func (u *UseCase) GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.PollOptionResponseWrapper, error) {
	myPollOption, err := u.PollOptionRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrModelNotFound()
		}

		return nil, customError.ErrModelGet(err, "PollOption")
	}

	return &presenter.PollOptionResponseWrapper{PollOption: myPollOption}, nil
}
