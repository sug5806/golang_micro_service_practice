package app

import (
	"golang_micro_service_practice/api/controller/hello"
	"golang_micro_service_practice/api/controller/repositories"
)

func url() {
	router.GET("/", hello.Hello)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
