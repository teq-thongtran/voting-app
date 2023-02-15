package payload

type CreateUserPollRequest struct {
	UserID int64 `json:"user_id" form:"user_id"`
	PollID int64 `json:"poll_id" form:"poll_id"`
}
