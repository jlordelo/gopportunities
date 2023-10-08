package main

import (
	"github.com/gin-gonic/gin"
	docs "github.com/go-project-name/docs"
	"github.com/jlordelo/gopportunities.git/config"
	"github.com/jlordelo/gopportunities.git/router"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	logger config.Logger
)

func main() {
	r := gin.Default()
	logger = *config.GetLogger("main")
	docs.SwaggerInfo.BasePath = "/api/v1"
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

	// initialize router
	router.Initialize()
}
