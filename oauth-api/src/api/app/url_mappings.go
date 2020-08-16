package app

import (
	"golang_micro_service_practice/api/controller/hello"
	"golang_micro_service_practice/oauth-api/src/api/controller/oauth"
)

func mapUrls() {
	router.GET("/hello", hello.Hello)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
