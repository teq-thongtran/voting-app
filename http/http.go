package http

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"myapp/auth"
	appses "myapp/http/app_session"
	"myapp/http/poll"
	pollopt "myapp/http/poll_option"
	"myapp/http/user"
	uspo "myapp/http/user_poll"
	usvo "myapp/http/user_vote"
	"myapp/repository"
	"myapp/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func NewHTTPHandler(useCase *usecase.UseCase, repo *repository.Repository) *echo.Echo {
	var (
		e         = echo.New()
		loggerCfg = middleware.DefaultLoggerConfig
	)

	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/health-check"
	}

	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOriginFunc: func(origin string) (bool, error) {
			return regexp.MatchString(
				`^https:\/\/(|[a-zA-Z0-9]+[a-zA-Z0-9-._]*[a-zA-Z0-9]+\.)teqnological.asia$`,
				origin,
			)
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))

	// APIs
	api := e.Group("/api")
	middlewares := auth.NewMiddlewareManager(repo.User)
	userApi := api.Group("/users", middlewares.RequiredAuth)
	pollApi := api.Group("/polls", middlewares.RequiredAuth)
	pollReferenceApi := api.Group("", middlewares.RequiredAuth)

	// Init groups APIs
	appses.Init(api.Group("/session"), useCase)
	user.Init(userApi.Group(""), useCase)
	poll.Init(pollApi.Group(""), useCase)
	pollopt.Init(pollReferenceApi.Group(""), useCase)
	uspo.Init(pollReferenceApi.Group(""), useCase)
	usvo.Init(pollReferenceApi.Group(""), useCase)
	return e
}
