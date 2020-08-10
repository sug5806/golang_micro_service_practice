package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang_micro_service_practice/api/clients/rest_client"
	"golang_micro_service_practice/api/domain/repositories"
	"golang_micro_service_practice/api/utils/errors"
	"golang_micro_service_practice/api/utils/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	CreateRepo(c)

	fmt.Println(response.Body.String())

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, "invalid json error", apiErr.ApiMessage())
	assert.EqualValues(t, http.StatusBadRequest, apiErr.ApiStatus())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()

	c := test_utils.GetMockContext(request, response)

	rest_client.FlushMockups()
	rest_client.AddMock(&rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.githubcom/v3/repos/docs"}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)

	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.ApiStatus())
	assert.EqualValues(t, "Requires authentication", apiErr.ApiMessage())

}

func TestCreateRepoNoError(t *testing.T) {
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

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)

}
