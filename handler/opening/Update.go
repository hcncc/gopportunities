package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Method PUT in endPoint opening",
	})
}
