package presenter

import (
	"myapp/model"
)

type UserVoteResponseWrapper struct {
	UserVote *model.UserVote `json:"user_vote"`
}

type ListUserVoteResponseWrapper struct {
	Poll      *model.Poll      `json:"poll"`
	UserVotes []model.UserVote `json:"user_votes"`
	Meta      interface{}      `json:"meta"`
}
