package auth

import (
	"errors"
	"github.com/labstack/echo/v4"
	"myapp/customError"
	"myapp/repository/user"
	"net/http"
	"strings"
)

type CustomMiddleware struct {
	userRepo user.Repository
}

func NewMiddlewareManager(userRepo user.Repository) *CustomMiddleware {
	return &CustomMiddleware{userRepo}
}

func extractTokenFromHeader(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	temp := strings.Split(bearerToken, " ")
	if temp[0] != "Bearer" || len(temp) < 2 || strings.TrimSpace(temp[1]) == "" {
		return "", errors.New("invalid token")
	}

	return temp[1], nil
}

func (a *CustomMiddleware) RequiredAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := extractTokenFromHeader(c.Request())
		if err != nil {
			return customError.ErrUnauthorized(err)
		}
		username, err := ValidateJWT(token)
		if err != nil {
			return customError.ErrUnauthorized(err)
		}

		if len(username) == 0 {
			return customError.ErrUnauthorized(err)
		}

		user, err := a.userRepo.GetByUsername(c.Request().Context(), username)
		if err != nil {
			return customError.ErrUnauthorized(err)
		}
		c.Set("user_id", user.ID)
		return next(c)
	}
}
