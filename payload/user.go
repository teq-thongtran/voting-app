package payload

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       int64   `json:"-"`
	Name     *string `json:"name"`
	Username *string `json:"username"`
}
