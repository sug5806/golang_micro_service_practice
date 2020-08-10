package repositories

import (
	"github.com/gin-gonic/gin"
	"golang_micro_service_practice/api/domain/repositories"
	"golang_micro_service_practice/api/services"
	"golang_micro_service_practice/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	createRepoRequest := repositories.CreateRepoRequest{}
	err := c.BindJSON(&createRepoRequest)
	if err != nil {
		apiErr := errors.NewApiBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}

	response, apiErr := services.RepositoryService.CreateRepo(createRepoRequest)
	if apiErr != nil {
		c.JSON(apiErr.ApiStatus(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func CreateRepos(c *gin.Context) {
	var createRepoRequests []repositories.CreateRepoRequest
	if err := c.BindJSON(&createRepoRequests); err != nil {
		apiErr := errors.NewApiBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(createRepoRequests)

	if err != nil {
		c.JSON(err.ApiStatus(), err)
		return
	}

	c.JSON(result.StatusCode, result)

}
