package graph

import (
	"fmt"
	"math"

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

func nestedMovies(actor models.Actor) []*model.Movie {
	var moviesGraph []*model.Movie
	for _, movie := range actor.Movies {
		var castGraph []*model.Actor
		database.DB.Joins("JOIN public.\"cast\" AS c ON c.actor_id = actors.id AND c.movie_id = ?", actor.ID).Find(&movie.Actors)
		for _, movieActor := range movie.Actors {
			movieActor := actorAssembler(movieActor)
			castGraph = append(castGraph, &movieActor)
		}
		movie := movieAssembler(movie)
		movie.Actors = castGraph
		moviesGraph = append(moviesGraph, &movie)
	}
	return moviesGraph
}

func generateInfo(page *int, count int64, pageLimit int) model.Info {
	totalPages := math.Ceil(float64(count) / float64(pageLimit))
	var next int
	var previous int
	if *page+1 <= int(totalPages) {
		next = *page + 1
	}
	if *page-1 > 0 {
		previous = *page - 1
	}
	var last int = int(totalPages)
	var countPointer int = int(count)
	info := &model.Info{
		Count:    &countPointer,
		Last:     &last,
		Next:     &next,
		Previous: &previous,
	}
	return *info
}
