package app

import (
	"github.com/gin-gonic/gin"
	"golang_micro_service_practice/api/log/option_a"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartApp() {
	option_a.Log.Info("about to map the urls", "step:1", "status:pending")
	url()
	option_a.Log.Info("urls successfully mapped", "step:2", "status:success")

	if err := router.Run(":7869"); err != nil {
		panic(err)
	}
}
