package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Method GET for show openigs in endPoint opening",
	})
}
