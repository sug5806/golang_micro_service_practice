package app

import (
	"golang_micro_service_practice/api/controller/hello"
	"golang_micro_service_practice/api/controller/repositories"
)

func url() {
	router.GET("/", hello.Hello)
	router.POST("/create_repo", repositories.CreateRepo)
	router.POST("/create_repos", repositories.CreateRepos)
}
