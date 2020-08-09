package github_provider

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang_micro_service_practice/api/clients/rest_client"
	"golang_micro_service_practice/api/domain/github"
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

func TestConstData(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

func TestGetAuthorizationHeader(t *testing.T) {
	headerAuthorizationFormat := getAuthorizationHeader("123")
	assert.EqualValues(t, "token 123", headerAuthorizationFormat)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Error:      errors.New("invalid rest_client response"),
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid rest_client response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	rest_client.FlushMockups()
	invalidJson, _ := os.Open("/invalidJson.json")
	defer invalidJson.Close()

	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       invalidJson,
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid github create repo response body", err.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message" : "1"`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Message)
}

func TestCreateRepoUnAuthroization(t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.githubcom/v3/repos/#create"}`)),
		},
		Error: nil,
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
	assert.EqualValues(t, "https://developer.githubcom/v3/repos/#create", err.DocumentationUrl)
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"`)),
		},
		Error: nil,
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "json unmarshal fail github create repo response", err.Message)
}

func TestCreateRepoNoError(t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id" : 123, "name": "golang_tutorial", "full_name": "golang_tuto"}`)),
		},
		Error: nil,
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.Id)
	assert.EqualValues(t, "golang_tutorial", response.Name)
	assert.EqualValues(t, "golang_tuto", response.FullName)

}
