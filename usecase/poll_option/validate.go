package pollopt

import (
	"context"
	"strconv"
	"strings"

	"myapp/customError"
	"myapp/model"
	"myapp/payload"
)

func (u *UseCase) validateCreate(ctx context.Context, req *payload.CreatePollOptionRequest) error {
	myPoll, err := u.PollRepo.GetByID(ctx, req.PollID)

	if err != nil {
		return customError.ErrModelGet(err, "Poll")
	}

	err = u.validatePoll(ctx, myPoll, ctx.Value("user_id").(int64))

	if err != nil {
		return err
	}

	req.OptionText = strings.TrimSpace(req.OptionText)
	err = u.validateText(req.OptionText)

	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdatePollOptionRequest) (*model.PollOption, error) {
	myPollOption, err := u.PollOptionRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, customError.ErrModelGet(err, "PollOption")
	}

	err = u.validatePollOption(myPollOption, ctx.Value("user_id").(int64))
	if err != nil {
		return nil, err
	}
	if req.OptionText != nil {
		*req.OptionText = strings.TrimSpace(*req.OptionText)
		err = u.validateText(*req.OptionText)

		if err != nil {
			return nil, err
		}
		myPollOption.OptionText = *req.OptionText
	}

	userID := ctx.Value("user_id").(int64)

	if myPollOption.UserID != userID {
		return nil, customError.ErrUnauthorized(nil)
	}

	return myPollOption, nil
}

func (u *UseCase) validateText(optionText string) error {
	if len(optionText) == 0 {
		return customError.ErrRequestInvalidParam("poll_option_text")
	}

	return nil
}

func (u *UseCase) validatePollOption(pollOption *model.PollOption, userID int64) error {
	if pollOption.UserID != userID {
		return customError.ErrGetByPolicty()
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
