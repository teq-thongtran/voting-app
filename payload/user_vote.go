package payload

type CreateUserVoteRequest struct {
	UserId       int64 `json:"user_id" form:"user_id"`
	PollOptionId int64 `json:"poll_option_id" form:"poll_option_id"`
}
