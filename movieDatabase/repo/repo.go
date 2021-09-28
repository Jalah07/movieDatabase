package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"movieDatabase/entities"
)

type Database struct {
	Movies []entities.MovieStruct
}


func AddMovie(m entities.MovieStruct) (Database, error) {
	file, err := ioutil.ReadFile("moviedb.json")
	if err != nil {
		fmt.Println(err)
	}
	m.SetId()

	sliceOfDb := Database{}

	err = json.Unmarshal(file, &sliceOfDb)
	if err != nil {
		fmt.Println(err)
	}
	sliceOfDb.Movies = append(sliceOfDb.Movies, m)

	data, err := json.MarshalIndent(sliceOfDb, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("moviedb.json", data, 0644)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	return sliceOfDb, nil
}
