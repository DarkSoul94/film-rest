package http

import (
	"github.com/DarkSoul94/golang-template/app"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, uc app.Usecase) {
	h := NewHandler(uc)

	router.POST("/film")
	router.GET("/film/:id")
	router.GET("/film")
}
