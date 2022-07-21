package json

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/DarkSoul94/film-rest/app"
	"github.com/DarkSoul94/film-rest/models"
)

type repoJson struct {
	Path string
}

func NewJsonRepo(path string) app.Repository {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0755)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(file)
	if len(bytes) == 0 {
		bytes, _ = json.Marshal(Films{
			Films: make([]models.Film, 0),
		})

		file.Write(bytes)
	}

	return &repoJson{
		Path: path,
	}
}

type Films struct {
	Films []models.Film `json:"films"`
}

func (r *repoJson) GetAllFilms() ([]models.Film, error) {
	var films Films

	file, err := os.OpenFile(r.Path, os.O_CREATE|os.O_RDONLY, 0755)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &films)
	if err != nil {
		return nil, err
	}

	return films.Films, nil
}

func (r *repoJson) AddFilm(film models.Film) error {
	var (
		films Films
		err   error
	)

	file, err := os.OpenFile(r.Path, os.O_CREATE|os.O_RDONLY, 0755)
	defer file.Close()
	if err != nil {
		return err
	}

	films.Films, err = r.GetAllFilms()
	if err != nil {
		return err
	}
	films.Films = append(films.Films, film)

	data, err := json.Marshal(films)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
