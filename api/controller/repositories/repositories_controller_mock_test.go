package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang_micro_service_practice/api/clients/rest_client"
	"golang_micro_service_practice/api/domain/repositories"
	"golang_micro_service_practice/api/services"
	"golang_micro_service_practice/api/utils/errors"
	"golang_micro_service_practice/api/utils/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	funcCreateRepo  func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func(clientId string, requests []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError)
)

type repoServiceMock struct{}

func init() {

}

func (r *repoServiceMock) CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo(clientId, request)
}
func (r *repoServiceMock) CreateRepos(clientId string, requests []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError) {
	return funcCreateRepos(clientId, requests)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Id:    321,
			Owner: "golang",
			Name:  "mocked service",
		}, nil
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()

	c := test_utils.GetMockContext(request, response)

	rest_client.FlushMockups()
	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})

	fmt.Println("sdafjfksd;k")

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, 321, result.Id)
	assert.EqualValues(t, "mocked service", result.Name)
	assert.EqualValues(t, "golang", result.Owner)

}

func TestCreateRepoErrorFromGithubMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewApiBadRequestError("invalid repository name")
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()

	c := test_utils.GetMockContext(request, response)

	rest_client.FlushMockups()
	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.ApiStatus())
	assert.EqualValues(t, "invalid repository name", apiErr.ApiMessage())

}
