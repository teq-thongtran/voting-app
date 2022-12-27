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
