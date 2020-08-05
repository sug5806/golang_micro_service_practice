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
		apiErr := errors.NewApiRequestError("invalid json error")
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
