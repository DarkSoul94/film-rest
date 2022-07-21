package http

import (
	"fmt"
	"net/http"

	"github.com/DarkSoul94/film-rest/app"
	"github.com/DarkSoul94/film-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	err = h.uc.CreateFilm(mFilm)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success"})
}

func (h *Handler) GetFilm(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error"})
		return
	}

	film, err := h.uc.GetFilm(id)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "empty"})
			return
		}
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "id": film.ID, "title": film.Title, "releaseDate": film.ReleaseDate})
}

func (h *Handler) GetFilmsList(c *gin.Context) {
	films, err := h.uc.GetAllFilms()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error"})
		return
	}

	if len(films) == 0 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "empty"})
		return
	}

	var outFilms []outFilm = make([]outFilm, 0)

	for _, film := range films {
		outFilms = append(outFilms, h.toOutFilm(film))
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "data": outFilms})
}
