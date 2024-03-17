package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/hcncc/gopportunities/handler/Opening"
)

func initializerRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		// Show one opening
		v1.GET("/opening/id", handler.ShowOpeningHandler)

		// Store opening
		v1.POST("/opening", handler.StoreOpeningHandler)

		//Update opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		// Delete opening
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		// FindMany opening
		v1.GET("/openings", handler.IndexOpeningHandler)
	}
}
