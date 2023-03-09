package httpserver

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/", IsPrimeNumber)
	return router
}
