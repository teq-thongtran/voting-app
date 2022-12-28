package payload

type CreatePollRequest struct {
	PollPolicy   string `json:"poll_policy" form:"poll_policy"`
	PollTitle    string `json:"poll_title" form:"poll_title"`
	PollVoteType string `json:"poll_vote_type" form:"poll_vote_type"`
}

type UpdatePollRequest struct {
	ID           int64   `json:"-"`
	PollPolicy   *string `json:"poll_policy" form:"poll_policy"`
	PollTitle    *string `json:"poll_title" form:"poll_title"`
	PollVoteType *string `json:"poll_vote_type" form:"poll_vote_type"`
}
