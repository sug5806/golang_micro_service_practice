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
	log.Log.Info("about to map the urls")
	url()
	log.Log.Info("urls successfully mapped")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
