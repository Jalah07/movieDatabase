package repo

import (
	"encoding/json"
	"io/ioutil"
	"movieProj/entities"
)

type movieStruct struct {
	Movies []entities.Movie
}

type Repo struct {
	Filename string
}


func NewRepo (filename string) Repo{
	return Repo {
		Filename : filename,
	}
}

func (r Repo) CreateNewMovie (mv entities.Movie) error {
	ms := movieStruct{}

	byteStruct, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	mv.SetId() 
	
	err = json.Unmarshal(byteStruct, &ms)
	if err != nil {
		return err
	}

	ms.Movies = append(ms.Movies, mv)
	byteSlice, err := json.MarshalIndent(ms, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, byteSlice, 0644)
	if err != nil {
		return err
	}
	return err
}