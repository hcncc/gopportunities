package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	router.Run(":8089") //listening port 8089
}
