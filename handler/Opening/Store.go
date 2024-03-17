package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoreOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Method POST in endPoint opening",
	})
}
