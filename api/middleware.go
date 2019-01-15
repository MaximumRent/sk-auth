package api

import "github.com/gin-gonic/gin"

func InitMiddleware(router *gin.Engine) {
	router.Use(mapMessageMiddleware)
}

func mapMessageMiddleware(context *gin.Context) {

	context.Next()
}
