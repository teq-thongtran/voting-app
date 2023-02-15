package payload

type CreateUserVoteRequest struct {
	UserID       int64 `json:"user_id" form:"user_id"`
	PollOptionID int64 `json:"poll_option_id" form:"poll_option_id"`
}
