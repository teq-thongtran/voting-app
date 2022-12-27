package customError

import (
	"fmt"
	"myapp/appError"
	"net/http"
)

func ErrModelGet(err error, modelName string) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10000",
		Message:   "Failed to get " + modelName,
		IsSentry:  true,
	}
}

func ErrModelCreate(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10001",
		Message:   "Failed to create record",
		IsSentry:  true,
	}
}

func ErrModelUpdate(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10002",
		Message:   "Failed to update record",
		IsSentry:  true,
	}
}

func ErrModelDelete(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10003",
		Message:   "Failed to delete record",
		IsSentry:  true,
	}
}

func ErrModelNotFound() appError.AppError {
	return appError.AppError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrRequestInvalidParam(param string) appError.AppError {
	return appError.AppError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10005",
		Message:   fmt.Sprintf("Invalid paramemter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrCommitTransaction(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "002",
		Message:   "Failed to commit transaction",
		IsSentry:  true,
	}
}

func ErrUnauthorized(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "000001",
		Message:   "Unauthorized!",
		IsSentry:  false,
	}
}

func ErrNoPermission() appError.AppError {
	return appError.AppError{
		Raw:       nil,
		HTTPCode:  http.StatusForbidden,
		ErrorCode: "000002",
		Message:   "No permission.",
		IsSentry:  false,
	}
}

func ErrInvalidParams(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "000003",
		Message:   "Invalid params.",
		IsSentry:  false,
	}
}

func ErrNotFound(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "000004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrGet(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000005",
		Message:   "Failed to get.",
		IsSentry:  true,
	}
}

func ErrCreate(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000006",
		Message:   "Failed to create.",
		IsSentry:  true,
	}
}

func ErrUpdate(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000007",
		Message:   "Failed to update.",
		IsSentry:  true,
	}
}

func ErrDelete(err error) appError.AppError {
	return appError.AppError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "000008",
		Message:   "Failed to delete.",
		IsSentry:  true,
	}
}
