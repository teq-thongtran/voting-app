package poll_option

import (
	"context"
	"strings"

	"myapp/customError"
	"myapp/model"
	"myapp/payload"
)

func (u *UseCase) validateCreate(ctx context.Context, req *payload.CreatePollOptionRequest) error {
	myPoll, err := u.PollRepo.GetByID(ctx, req.PollId)

	if err != nil {
		return err
	}

	err = u.validatePoll(myPoll, ctx.Value("user_id").(int64))

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

	userId := ctx.Value("user_id").(int64)

	if myPollOption.UserId != userId {
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

func (u *UseCase) validatePollOption(pollOption *model.PollOption, userId int64) error {
	if pollOption.UserId != userId {
		return customError.ErrGetByPolicty()
	}
	return nil
}

func (u *UseCase) validatePoll(poll *model.Poll, userId int64) error {
	if poll.UserId != userId {
		return customError.ErrGetByPolicty()
	}
	return nil
}