package http

import (
	"github.com/DarkSoul94/film-rest/app"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, uc app.Usecase) {
	h := NewHandler(uc)

	router.POST("/film", h.CreateFilm)
	router.GET("/film/:id", h.GetFilm)
	router.GET("/film", h.GetFilmsList)
}
