package usvo

import (
	"strconv"

	"myapp/appError"
	"myapp/customError"
	"myapp/payload"
	"myapp/presenter"
	"myapp/teq"
	"myapp/usecase"

	"github.com/labstack/echo/v4"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}
	group.POST("/polls/:poll_id/user_votes", r.Create)
	group.GET("/polls/:poll_id/user_votes", r.GetList)
	group.DELETE("/user_votes/:id", r.Delete)
}

func (r *Route) Create(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		resp *presenter.UserVoteResponseWrapper
		req  = payload.CreateUserVoteRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.UserVote.Create(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(appError.AppError))
	}

	return teq.Response.Success(c, resp)
}

func (r *Route) Delete(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	err = r.UseCase.UserVote.Delete(ctx, &payload.DeleteRequest{ID: id})
	if err != nil {
		return teq.Response.Error(c, err.(appError.AppError))
	}

	return teq.Response.Success(c, nil)
}

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx       = &teq.CustomEchoContext{Context: c}
		req       = payload.GetListRequest{}
		resp      *presenter.ListUserVoteResponseWrapper
		pollIDStr = c.Param("poll_id")
	)

	pollID, err := strconv.ParseInt(pollIDStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.UserVote.GetList(ctx, &req, pollID)
	if err != nil {
		return teq.Response.Error(c, err.(appError.AppError))
	}

	return teq.Response.Success(c, resp)
}
