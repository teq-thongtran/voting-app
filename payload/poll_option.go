package payload

type CreatePollOptionRequest struct {
	UserId     int64  `json:"user_id" form:"user_id"`
	PollId     int64  `json:"poll_id" form:"poll_id"`
	OptionText string `json:"option_text" form:"option_text"`
}

type UpdatePollOptionRequest struct {
	ID         int64   `json:"-"`
	OptionText *string `json:"option_text" form:"option_text"`
}
