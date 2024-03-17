package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hcncc/gopportunities/handler"
)

func initializerRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		// Show one opening
		v1.GET("/opening/id", handler.Opening.Show.ShowOpeningHandler)

		// Store opening
		v1.POST("/opening", handler.StoreOpeningHandler)

		//Update opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.GET("/openings", handler.IndexOpeningHandler)
	}
}
