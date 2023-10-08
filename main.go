package main

import (
	"github.com/jlordelo/gopportunities.git/config"
	"github.com/jlordelo/gopportunities.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// initialize router
	router.Initialize()
}
