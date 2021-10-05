package service

import (
	"errors"
	"movieProj/entities"
	"movieProj/repo"
)

type Service struct {
	Repo repo.Repo
}

func NewService(r repo.Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) CreateNewMovie(mv entities.Movie) error {
	mv.SetId()
	if mv.Rating >= 0 && mv.Rating <= 10 {
		s.Repo.CreateNewMovie(mv)
	}
	return errors.New("invaild rating")
}

func (s Service) FindById(id string) (entities.Movie, error) {
	movie, err := s.Repo.FindById(id)
	if err != nil {
		return movie, nil
	}
	return movie, nil
}

func (s Service) GetMovies() (repo.MovieStruct, error) {
	mv, err := s.Repo.GetMovies()
	if err != nil {
		return mv, err
	}
	return mv, nil
}

func (s Service) DeleteMovie(id string) (entities.Movie, error) {
	movie, err := s.Repo.FindById(id)
	if err != nil {
		return movie, nil
	}
	return movie, nil
}
