package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	router := gin.Default()
	// Initialize routes
	initializeRoutes(router)

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// Run the server
	router.Run(":8080")
}
