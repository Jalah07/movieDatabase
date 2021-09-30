package handlers

import (
/* 	"encoding/json"
	"movieProj/entities"
	"movieProj/service"
	"net/http" */

	"github.com/gorilla/mux"
)


type Struct struct {

}



func ConfigureRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

 	r.HandleFunc("/movies", handler.PostMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", handler.GetMovieHandler).Methods("GET")
	return r
}
