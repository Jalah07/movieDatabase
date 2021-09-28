package handlers

import (
	"encoding/json"
	"fmt"
	"movieDatabase/entities"
	"movieDatabase/repo"
	"net/http"
)


func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.MovieStruct{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	db, err := repo.AddMovie(mv) // Calls addMovie inorder to post movie
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
