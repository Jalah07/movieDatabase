package entities

import "github.com/google/uuid"

type MovieStruct struct {
	Id          string
	Title       string
	Genre       []string
	Description string
	Director    string
	Actors      []string
	Rating      float64
}

func (mv *MovieStruct) SetId() {
	mv.Id = uuid.New().String()
}
