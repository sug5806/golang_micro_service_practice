package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	ApiStatus() int
	ApiMessage() string
	ApiError() string
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (a *apiError) ApiStatus() int {
	return a.Code
}

func (a *apiError) ApiMessage() string {
	return a.Message
}

func (a *apiError) ApiError() string {
	return a.Error
}

func NewApiErrFromBytes(body []byte) (ApiError, error) {
	var result apiError

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}

func NewApiError(status int, message string) ApiError {
	return &apiError{
		Code:    status,
		Message: message,
	}
}

func NewApiRequestError(message string) ApiError {
	return &apiError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewApiResponseError(message string) ApiError {
	return &apiError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewApiBadRequestError(message string) ApiError {
	return &apiError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}
