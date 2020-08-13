package app

import (
	"github.com/gin-gonic/gin"
	"golang_micro_service_practice/api/log"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Log.Info("about to map the urls", "step:1", "status:pending")
	url()
	log.Log.Info("urls successfully mapped", "step:2", "status:success")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
