package services

import (
	"golang_micro_service_practice/api/config"
	"golang_micro_service_practice/api/domain/github"
	"golang_micro_service_practice/api/domain/repositories"
	"golang_micro_service_practice/api/provider/github_provider"
	"golang_micro_service_practice/api/utils/errors"
	"net/http"
	"strings"
	"sync"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(requests []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError)
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

func (r *repoService) CreateRepos(requests []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	var wg sync.WaitGroup
	defer close(output)

	go r.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		r.CreateRepo(current)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCreation := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreation++
		}
	}

	if successCreation == 0 {
		result.StatusCode = result.Results[0].Error.ApiStatus()
	} else if successCreation == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return &result, nil
}

func (r *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)

		wg.Done()
	}
	output <- results
}

func (r *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := r.CreateRepo(input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}
}
