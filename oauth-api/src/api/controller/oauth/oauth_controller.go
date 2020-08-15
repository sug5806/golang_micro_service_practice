package oauth

import (
	"github.com/gin-gonic/gin"
	"golang_micro_service_practice/api/utils/errors"
	"golang_micro_service_practice/oauth-api/src/api/domain/oauth"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewApiBadRequestError("invalid json baody")
		c.JSON(apiErr.ApiStatus(), apiErr)
		return
	}

}
