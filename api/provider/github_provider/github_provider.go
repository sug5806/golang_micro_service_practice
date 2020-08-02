package github_provider

import (
	"encoding/json"
	"fmt"
	"golang_micro_service_practice/api/clients/rest_client"
	"golang_micro_service_practice/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, requestBody github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := rest_client.Post(urlCreateRepo, requestBody, headers)
	if err != nil {
		log.Printf("Error when trying to create new repo in github : %s \n\n", err.Error())
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid github create repo response body",
		}
	}

	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "json unmarshal fail github create repo response",
		}
	}

	return &result, nil
}
