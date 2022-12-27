package payload

type CreatePollRequest struct {
	PollPolicy   int8   `json:"poll_policy"`
	PollTitle    string `json:"poll_title"`
	PollVoteType int8   `json:"poll_vote_type"`
}

type UpdatePollRequest struct {
	ID           int64   `json:"-"`
	PollPolicy   int8    `json:"poll_policy"`
	PollTitle    *string `json:"poll_title"`
	PollVoteType int8    `json:"poll_vote_type"`
}
