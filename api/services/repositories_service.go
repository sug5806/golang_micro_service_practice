package services

import (
	"golang_micro_service_practice/api/config"
	"golang_micro_service_practice/api/domain/github"
	"golang_micro_service_practice/api/domain/repositories"
	"golang_micro_service_practice/api/provider/github_provider"
	"golang_micro_service_practice/api/utils/errors"
	"strings"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (r *repoService) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	name := strings.TrimSpace(request.Name)
	if name == "" {
		return nil, errors.NewApiBadRequestError("invalid repository name")
	}

	githubRequest := github.CreateRepoRequest{
		Name:            request.Name,
		Description:     request.Description,
		Private:         false,
		LicenseTemplate: "mit",
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), githubRequest)

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	return &repositories.CreateRepoResponse{
		Id:    uint64(response.Id),
		Owner: response.Owner.Login,
		Name:  response.Name,
	}, nil
}
