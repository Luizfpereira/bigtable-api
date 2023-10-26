package router

import (
	"bigtable_api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(climateHandler *handlers.ClimateHandler) *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"}) })
	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })

	read := router.Group("/read")
	read.GET("/climate-data", climateHandler.ReadClimateData)
	return router
}
