package uspo

import (
	"context"

	"myapp/customError"
	"myapp/model"
	"myapp/payload"
)

func (u *UseCase) validateCreate(ctx context.Context, req *payload.CreateUserPollRequest) error {
	myPoll, err := u.PollRepo.GetByID(ctx, req.PollID)

	if err != nil {
		return customError.ErrModelGet(err, "Poll")
	}

	err = u.validatePoll(myPoll, ctx.Value("user_id").(int64))

	if err != nil {
		return err
	}

	_, err = u.UserRepo.GetByID(ctx, req.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) validatePoll(poll *model.Poll, userID int64) error {
	if poll.UserID != userID {
		return customError.ErrGetByPolicty()
	}
	return nil
}
