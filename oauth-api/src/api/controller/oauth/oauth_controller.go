package oauth

import (
	"github.com/gin-gonic/gin"
	"golang_micro_service_practice/api/utils/errors"
	"golang_micro_service_practice/oauth-api/src/api/domain/oauth"
	"golang_micro_service_practice/oauth-api/src/api/services"
	"net/http"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewApiBadRequestError("invalid json baody")
		c.JSON(apiErr.ApiStatus(), apiErr)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.ApiStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, token)

}

func GetAccessToken(c *gin.Context) {
	tokenId := c.Param("token_id")

	token, err := services.OauthService.GetAccessToken(tokenId)

	if err != nil {
		c.JSON(err.ApiStatus(), err)
		return
	}

	c.JSON(http.StatusOK, token)

}
