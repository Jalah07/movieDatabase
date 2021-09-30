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
		Repo : r,
	}
}

func (s Service) CreateNewMovie(mv entities.Movie) error {
	if mv.Rating >= 0 && mv.Rating <= 10 {
		s.Repo.CreateNewMovie(mv)
	}
	return errors.New("invaild rating")
} 