package graph

import (
	"fmt"

	"github.com/kuro-vale/kuro-movies-api/graph/model"
	"github.com/kuro-vale/kuro-movies-api/models"
)

func movieAssembler(movie models.Movie) model.Movie {
	id := fmt.Sprint(movie.ID)
	movieGraph := model.Movie{
		ID:       &id,
		Title:    &movie.Title,
		Genre:    &movie.Genre,
		Price:    &movie.Price,
		Director: &movie.Director,
		Producer: &movie.Producer,
	}
	return movieGraph
}

func actorAssembler(actor models.Actor) model.Actor {
	var agePointer int = int(actor.Age)
	id := fmt.Sprint(actor.ID)
	actorGraph := model.Actor{
		ID:     &id,
		Name:   &actor.Name,
		Age:    &agePointer,
		Gender: &actor.Gender,
	}
	return actorGraph
}
