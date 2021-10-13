package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"movieProj/entities"
	"movieProj/repo"
	"net/http"
)

type Service interface {
	CreateNewMovie(mv entities.Movie) error
	GetMovies() (repo.MovieStruct, error)
	FindById(id string) (entities.Movie, error)
	DeleteMovie(id string) (err error)
	UpdateMovie(mv entities.Movie, id string) (err error)
}

type MovieHandler struct {
	Svc Service
}

func NewMovieHandler(s Service) MovieHandler {
	return MovieHandler{
		Svc: s,
	}
}

func (mh MovieHandler) PostMovieHandler(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = mh.Svc.CreateNewMovie(mv)

	if err != nil {
		switch err.Error() {
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusBadRequest)

		case "movie does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)
		}

	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (mh MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	mvID, err := mh.Svc.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	movie, err := json.MarshalIndent(mvID, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(movie)
}

func (mh MovieHandler) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	myDb, err := mh.Svc.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	movieDb, err := json.MarshalIndent(myDb, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(movieDb)
}

func (mh MovieHandler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := mh.Svc.DeleteMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (mh MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Check this error code
	}

	err = mh.Svc.UpdateMovie(mv, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
