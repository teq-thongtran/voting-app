package pollopt

import (
	"myapp/appError"
	"myapp/customError"
	"myapp/payload"
	"myapp/presenter"
	"myapp/teq"
	"myapp/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}
	group.POST("/polls/:poll_id/poll_options", r.Create)
	group.GET("/polls/:poll_id/poll_options", r.GetList)
	group.PUT("/poll_options/:id", r.Update)
	group.DELETE("/poll_options/:id", r.Delete)
}

func (r *Route) Create(c echo.Context) error {
	var (
		ctx       = &teq.CustomEchoContext{Context: c}
		resp      *presenter.PollOptionResponseWrapper
		req       = payload.CreatePollOptionRequest{}
		pollIdStr = c.Param("poll_id")
	)
	pollId, err := strconv.ParseInt(pollIdStr, 10, 64)

	if err != nil {
		return err
	}

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	req.PollId = pollId
	resp, err = r.UseCase.PollOption.Create(ctx, &req)
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

	err = r.UseCase.PollOption.Delete(ctx, &payload.DeleteRequest{ID: id})
	if err != nil {
		return teq.Response.Error(c, err.(appError.AppError))
	}

	return teq.Response.Success(c, nil)
}

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx       = &teq.CustomEchoContext{Context: c}
		req       = payload.GetListRequest{}
		resp      *presenter.ListPollOptionResponseWrapper
		pollIdStr = c.Param("poll_id")
	)

	pollId, err := strconv.ParseInt(pollIdStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.PollOption.GetList(ctx, &req, pollId)
	if err != nil {
		return teq.Response.Error(c, err.(appError.AppError))
	}

	return teq.Response.Success(c, resp)
}

func (r *Route) Update(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.PollOptionResponseWrapper
	)

	PollOptionId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	req := payload.UpdatePollOptionRequest{
		ID: PollOptionId,
	}

	if err = c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.PollOption.Update(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(appError.AppError))
	}

	return teq.Response.Success(c, resp)
}
