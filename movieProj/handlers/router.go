package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostMovieHandler).Methods("POST")
	r.HandleFunc("/movie/{Id}", handler.GetMovieHandler).Methods("GET")
	r.HandleFunc("/movie", handler.GetMoviesHandler).Methods("GET")
	r.HandleFunc("/movie/{Id}", handler.DeleteMovieHandler).Methods("DELETE")
	r.HandleFunc("/movie/{Id}", handler.UpdateMovieHandler).Methods("PUT")

	return r
}
