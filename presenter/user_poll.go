package presenter

import (
	"myapp/model"
)

type UserPollResponseWrapper struct {
	UserPoll *model.UserPoll `json:"user_poll"`
}

type ListUserPollResponseWrapper struct {
	Poll      *model.Poll      `json:"poll"`
	UserPolls []model.UserPoll `json:"user_polls"`
	Meta      interface{}      `json:"meta"`
}
