package repo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"movieProj/entities"
)

type MovieStruct struct { // m
	Movies []entities.Movie
}

type Repo struct {
	Filename string
}

func NewRepo(filename string) Repo {
	return Repo{
		Filename: filename,
	}
}

func (r Repo) CreateNewMovie(mv entities.Movie) error {
	ms := MovieStruct{}

	byteStruct, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

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
	return nil
}

func (r Repo) FindById(id string) (entities.Movie, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return entities.Movie{}, errors.New("could not read file")
	}

	movies := MovieStruct{}
	err = json.Unmarshal(file, &movies)
	if err != nil {
		return entities.Movie{}, errors.New("failed to unmarshal content")
	}
	moviedb := entities.Movie{}

	for _, movie := range movies.Movies {
		if movie.Id == id {
			moviedb = movie
			return moviedb, nil
		}
	}
	return moviedb, nil

}

func (r Repo) GetMovies() (MovieStruct, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return MovieStruct{}, err
	}

	movies := MovieStruct{}
	err = json.Unmarshal(file, &movies)
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (r Repo) DeleteMovie(id string) error {
	movies := MovieStruct{}
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return errors.New("could not read file")
	}

	err = json.Unmarshal(file, &movies)
	if err != nil {
		return err
	}

	for index, movie := range movies.Movies {
		if movie.Id == id {
			movies.Movies = append(movies.Movies[:index], movies.Movies[index+1:]...)
			output, err := json.MarshalIndent(&movies, "", " ")
			if err != nil {
				return err
			}
			ioutil.WriteFile(r.Filename, output, 0644)
			return nil
		}
	}
	return nil
}

func (r Repo) UpdateMovie(mv entities.Movie, id string) error {
	movies := MovieStruct{}
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &movies)
	if err != nil {
		return err
	}

	for index, movie := range movies.Movies {
		if movie.Id == id {
			movies.Movies = append(movies.Movies[:index], movies.Movies[index+1:]...)
			movie.Id = id
			movies.Movies = append(movies.Movies, mv)
		}
	}
	output, err := json.MarshalIndent(&movies, "", " ")
	if err != nil {
		return err
	}
	ioutil.WriteFile(r.Filename, output, 0644)
	return err

}
