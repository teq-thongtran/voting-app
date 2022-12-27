package appError

import (
	"errors"
	"fmt"
	"strings"
)

type AppError struct {
	Raw       error
	ErrorCode string
	HTTPCode  int
	Message   string
	IsSentry  bool
}

func (e AppError) Error() string {
	if e.Raw != nil {
		fmt.Println(e.Raw)
	}

	return e.Message
}

func (e AppError) Is(target error) bool {
	if e.Raw != nil {
		return errors.Is(e.Raw, target)
	}

	return strings.Contains(e.Error(), target.Error())
}

func NewError(err error, httpCode int, errCode string, message string, isSentry bool) AppError {
	return AppError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
		IsSentry:  isSentry,
	}
}
