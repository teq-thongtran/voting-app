package auth

import (
	"errors"
	"myapp/repository/user"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
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
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Header invalid!",
			})
		}
		username, err := ValidateJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Validate token get error",
			})
		}

		if len(username) == 0 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Token invalid!",
			})
		}

		user, err := a.userRepo.GetByUsername(c.Request().Context(), username)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Token not match any users",
			})
		}
		c.Set("user_id", user.ID)
		return next(c)
	}
}
