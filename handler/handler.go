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

func StoreOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Method POST in endPoint opening",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Method GET in endPoint opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Method PUT in endPoint opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Method DELETE in endPoint opening",
	})
}
