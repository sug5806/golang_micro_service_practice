package oauth

import (
	"golang_micro_service_practice/api/utils/errors"
	"strings"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a AccessTokenRequest) Validate() errors.ApiError {
	a.Username = strings.TrimSpace(a.Username)
	if a.Username == "" {
		return errors.NewApiBadRequestError("invalid username")
	}

	a.Password = strings.TrimSpace(a.Password)
	if a.Password == "" {
		return errors.NewApiBadRequestError("invalid password")
	}

	return nil
}
