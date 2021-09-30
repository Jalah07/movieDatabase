package handlers

import (
	"encoding/json"
	"fmt"

	"movieProj/entities"
	"movieProj/service"
	"net/http"

	"github.com/gorilla/mux"
)



type MovieHandler struct {
	Svc service.Service
}



func NewMovieHandler(s service.Service) MovieHandler {
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



func (mh MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	id := params["Id"]
	fmt.Println(id)
/* 	movie := []entities.Movie{}
	//fmt.Println(movie)
	for _, item := range movie {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			
			
			return 
		} 
		fmt.Println(item.Id)
	}
	//fmt.Println(params)
	json.NewEncoder(w).Encode(&entities.Movie{}) */
	//w.WriteHeader(http.StatusCreated)
	//w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(item))
	
} 