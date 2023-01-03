package presenter

import (
	"myapp/model"
)

type PollResponseWrapper struct {
	Poll *model.Poll `json:"poll"`
}

type ListPollResponseWrapper struct {
	Polls []model.Poll `json:"polls"`
	Meta  interface{}  `json:"meta"`
}

type PollDetailResponseWrapper struct {
	Poll        *model.Poll        `json:"poll"`
	PollOptions []model.PollOption `json:"poll_options"`
	UserVotes   []model.UserVote   `json:"user_votes"`
}
