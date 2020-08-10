package repositories

import (
	"golang_micro_service_practice/api/utils/errors"
	"strings"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRepoResponse struct {
	Id    uint64 `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int                        `json:"status_code"`
	Results    []CreateRepositoriesResult `json:"results"`
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"response"`
	Error    errors.ApiError     `json:"error"`
}

func (c *CreateRepoRequest) Validate() errors.ApiError {
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		return errors.NewApiBadRequestError("invalid repository name")
	}
	return nil
}
