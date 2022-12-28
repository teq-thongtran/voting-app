package presenter

import (
	"myapp/model"
)

type PollOptionResponseWrapper struct {
	PollOption *model.PollOption `json:"poll_options"`
}

type ListPollOptionResponseWrapper struct {
	PollOptions []model.PollOption `json:"poll_options"`
	Meta        interface{}        `json:"meta"`
}
