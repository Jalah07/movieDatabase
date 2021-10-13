package service

import (
	"errors"
	"movieProj/entities"
	"movieProj/repo"
)

type Repository interface {
	CreateNewMovie(mv entities.Movie) error
	GetMovies() (repo.MovieStruct, error)
	FindById(id string) (entities.Movie, error)
	DeleteMovie(id string) (err error)
	UpdateMovie(mv entities.Movie, id string) (err error)
}

type Service struct {
	Repo Repository
}

func NewService(r Repository) Service {
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

// return errors.Wrap(err, "some information")
func (s Service) FindById(id string) (entities.Movie, error) {
	movie, err := s.Repo.FindById(id)
	if err != nil {
		return movie, err
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

func (s Service) DeleteMovie(id string) (err error) {
	if err != nil {
		return err
	}
	return s.Repo.DeleteMovie(id)
}

func (s Service) UpdateMovie(mv entities.Movie, id string) (err error) {
	err = s.Repo.UpdateMovie(mv, id)
	if err != nil {
		return err
	}
	return nil
}
