package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/graph/generated"
	"github.com/kuro-vale/kuro-movies-api/graph/model"
	"github.com/kuro-vale/kuro-movies-api/models"
)

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var response []*model.Movie
	var movies []models.Movie
	database.DB.Preload("Actors").Find(&movies)
	for _, movie := range movies {
		var actors []*model.Actor
		for _, actor := range movie.Actors {
			var actorMovies []*model.Movie
			database.DB.Joins("JOIN public.\"cast\" AS c ON c.actor_id = ? AND c.movie_id = movies.id", actor.ID).Find(&actor.Movies)
			for _, actorMovie := range actor.Movies {
				actorMovie := movieAssembler(actorMovie)
				actorMovies = append(actorMovies, &actorMovie)
			}
			actor := actorAssembler(actor)
			actor.Movies = actorMovies
			actors = append(actors, &actor)
		}
		movie := movieAssembler(movie)
		movie.Actors = actors
		response = append(response, &movie)
	}
	return response, nil
}

func (r *queryResolver) Actors(ctx context.Context) ([]*model.Actor, error) {
	var response []*model.Actor
	var actors []models.Actor
	database.DB.Preload("Movies").Find(&actors)
	for _, actor := range actors {
		var movies []*model.Movie
		for _, movie := range actor.Movies {
			var cast []*model.Actor
			database.DB.Joins("JOIN public.\"cast\" AS c ON c.actor_id = actors.id AND c.movie_id = ?", actor.ID).Find(&movie.Actors)
			for _, movieActor := range movie.Actors {
				movieActor := actorAssembler(movieActor)
				cast = append(cast, &movieActor)
			}
			movie := movieAssembler(movie)
			movie.Actors = cast
			movies = append(movies, &movie)
		}
		actor := actorAssembler(actor)
		actor.Movies = movies
		response = append(response, &actor)
	}
	return response, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func movieAssembler(movie models.Movie) model.Movie {
	movieGraph := model.Movie{
		ID:       fmt.Sprint(movie.ID),
		Title:    movie.Title,
		Genre:    movie.Genre,
		Price:    movie.Price,
		Director: movie.Director,
		Producer: movie.Producer,
	}
	return movieGraph
}

func actorAssembler(actor models.Actor) model.Actor {
	actorGraph := model.Actor{
		ID:     fmt.Sprint(actor.ID),
		Name:   actor.Name,
		Age:    int(actor.Age),
		Gender: actor.Gender,
	}
	return actorGraph
}
