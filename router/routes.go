package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializerRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		// Show one opening
		v1.GET("/opening/id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Method GET in endpoint opening",
			})
		})

		// Store opening
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "Method POST in endPoint opening",
			})
		})

		//Update opening
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Method PUT in endPoint opening",
			})
		})

		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Method DELETE in endpoint opening",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Method GET for show openings",
			})
		})
	}
}
