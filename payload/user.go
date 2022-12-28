package payload

type CreateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UpdateUserRequest struct {
	ID       int64   `json:"-"`
	Name     *string `json:"name" form:"name"`
	Username *string `json:"username" form:"username"`
}
