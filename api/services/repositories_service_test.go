package services

import (
	"github.com/stretchr/testify/assert"
	"golang_micro_service_practice/api/clients/rest_client"
	"golang_micro_service_practice/api/domain/repositories"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestRepoService_CreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	response, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid repository name", err.ApiMessage())

}

func TestRepoService_CreateRepoErrorFromGithub(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.githubcom/v3/repos/docs"}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "test",
	}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.ApiStatus())
	assert.EqualValues(t, "Requires authentication", err.ApiMessage())

}

func TestRepoService_CreateRepoNoError(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "sug5806"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "",
	}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "sug5806", result.Owner)
}
