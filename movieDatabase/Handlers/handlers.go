package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"movieDatabase/entities"
	"movieDatabase/repo"
)

func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.MovieStruct{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	mv.SetId()

	db := repo.MVStruct{}
	
	jsonBytes, err := json.MarshalIndent(db, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("moviedb.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
