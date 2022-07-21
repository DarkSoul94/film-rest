package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DarkSoul94/film-rest/app"
	"github.com/DarkSoul94/film-rest/models"
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	uc app.Usecase
}

// NewHandler ...
func NewHandler(uc app.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) CreateFilm(c *gin.Context) {
	var (
		film  newFilm
		mFilm models.Film
		err   error
	)

	err = c.BindJSON(&film)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "bodyInvalid"})
		return
	}

	mFilm, err = h.toModelFilm(film)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "bodyInvalid"})
		return
	}

	err = h.uc.CreateFilm(context.Background(), mFilm)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success"})
}

func (h *Handler) GetFilm(c *gin.Context) {

}

func (h *Handler) GetFilmsList(c *gin.Context) {

}
