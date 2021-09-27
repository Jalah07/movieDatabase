package entities

import "github.com/google/uuid"

type MovieStruct struct {
	id string
	title string
	genre string
	description string
	director string
	actors []string
	rating float64
}

func (mv *MovieStruct) SetId() {
	mv.id = uuid.New().String()
}