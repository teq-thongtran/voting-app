package usvo

import (
	"context"
	"strconv"

	"myapp/customError"
	"myapp/model"
	"myapp/payload"
)

func (u *UseCase) validateCreate(ctx context.Context, req *payload.CreateUserVoteRequest) error {
	myPollOption, err := u.PollOptionRepo.GetByID(ctx, req.PollOptionID)

	if err != nil {
		return customError.ErrModelGet(err, "PollOption")
	}

	myPoll, err := u.PollRepo.GetByID(ctx, myPollOption.PollID)

	err = u.validatePoll(ctx, myPoll, ctx.Value("user_id").(int64))

	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) validatePoll(ctx context.Context, poll *model.Poll, userID int64) error {
	ids, err := u.PollRepo.GetListPollID(ctx)
	if err != nil {
		return customError.ErrModelGet(err, "user_polls")
	}

	is_exists := contains(ids, strconv.FormatInt(poll.ID, 10))
	if poll.UserID == userID || is_exists || poll.PollPolicy == "public" {
		return nil
	}
	return customError.ErrGetByPolicty()
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
