package payload

import "strings"

type GetByIDRequest struct {
	ID int64 `json:"-"`
}

var orderBy = []string{"id", "created_by", "updated_by"}

type GetListRequest struct {
	Page    int    `json:"page" query:"page"`
	Limit   int    `json:"limit" query:"limit"`
	OrderBy string `json:"order_by,omitempty" query:"order_by"`
}

func (g *GetListRequest) Format() {
	g.OrderBy = strings.ToLower(strings.TrimSpace(g.OrderBy))

	for i := range orderBy {
		if g.OrderBy == orderBy[i] {
			return
		}
	}

	g.OrderBy = ""
}

type DeleteRequest struct {
	ID int64 `json:"-"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
