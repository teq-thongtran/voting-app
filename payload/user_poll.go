package payload

type CreateUserPollRequest struct {
	UserId int64 `json:"user_id" form:"user_id"`
	PollId int64 `json:"poll_id" form:"poll_id"`
}
