package poll

import (
	"context"
	"strings"

	"myapp/customError"
	"myapp/model"
	"myapp/payload"
)

func (u *UseCase) validateCreate(req *payload.CreatePollRequest) error {

	req.PollTitle = strings.TrimSpace(req.PollTitle)
	err := u.validateTitle(req.PollTitle)

	if err != nil {
		return err
	}

	req.PollPolicy = strings.TrimSpace(req.PollPolicy)
	req.PollPolicy = strings.ToLower(req.PollPolicy)
	err = u.validatePolicy(req.PollPolicy)

	if err != nil {
		return err
	}

	req.PollVoteType = strings.TrimSpace(req.PollVoteType)
	req.PollVoteType = strings.ToLower(req.PollVoteType)
	err = u.validateVoteType(req.PollVoteType)

	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdatePollRequest) (*model.Poll, error) {
	myPoll, err := u.PollRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, customError.ErrModelGet(err, "Poll")
	}

	err = u.validatePoll(myPoll, ctx.Value("user_id").(int64))
	if err != nil {
		return nil, err
	}

	if req.PollTitle != nil {
		*req.PollTitle = strings.TrimSpace(*req.PollTitle)
		err := u.validateTitle(*req.PollTitle)
		if err != nil {
			return nil, err
		}

		myPoll.PollTitle = *req.PollTitle
	}

	if req.PollPolicy != nil {
		*req.PollPolicy = strings.TrimSpace(*req.PollPolicy)
		*req.PollPolicy = strings.ToLower(*req.PollPolicy)
		err := u.validatePolicy(*req.PollPolicy)
		if err != nil {
			return nil, err
		}

		myPoll.PollPolicy = *req.PollPolicy
	}

	if req.PollVoteType != nil {
		*req.PollVoteType = strings.TrimSpace(*req.PollVoteType)
		*req.PollVoteType = strings.ToLower(*req.PollVoteType)
		err := u.validateVoteType(*req.PollVoteType)
		if err != nil {
			return nil, err
		}

		myPoll.PollVoteType = *req.PollVoteType
	}

	return myPoll, nil
}

func (u *UseCase) validateTitle(title string) error {
	if len(title) == 0 {
		return customError.ErrRequestInvalidParam("poll_title")
	}

	return nil
}

func (u *UseCase) validateVoteType(voteType string) error {
	if len(voteType) == 0 {
		return customError.ErrRequestInvalidParam("poll_vote_type")
	}

	if voteType != "single" && voteType != "multi" {
		return customError.ErrRequestInvalidParam("poll_vote_type")
	}

	return nil
}

func (u *UseCase) validatePolicy(policy string) error {
	if len(policy) == 0 {
		return customError.ErrRequestInvalidParam("poll_policy")
	}

	if policy != "public" && policy != "private" {
		return customError.ErrRequestInvalidParam("poll_policy")
	}

	return nil
}

func (u *UseCase) validatePoll(poll *model.Poll, userId int64) error {
	if poll.UserId != userId {
		return customError.ErrGetByPolicty()
	}
	return nil
}
