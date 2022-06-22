package graph

import (
	"fmt"

	"github.com/kuro-vale/kuro-movies-api/database"
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

func nestedActors(movie models.Movie) []*model.Actor {
	var actorsGraph []*model.Actor
	for _, actor := range movie.Actors {
		var actorMoviesGraph []*model.Movie
		database.DB.Joins("JOIN public.\"cast\" AS c ON c.actor_id = ? AND c.movie_id = movies.id", actor.ID).Find(&actor.Movies)
		for _, actorMovie := range actor.Movies {
			actorMovie := movieAssembler(actorMovie)
			actorMoviesGraph = append(actorMoviesGraph, &actorMovie)
		}
		actor := actorAssembler(actor)
		actor.Movies = actorMoviesGraph
		actorsGraph = append(actorsGraph, &actor)
	}
	return actorsGraph
}
