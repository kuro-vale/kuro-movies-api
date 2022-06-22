package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math"

	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/graph/generated"
	"github.com/kuro-vale/kuro-movies-api/graph/model"
	"github.com/kuro-vale/kuro-movies-api/models"
)

func (r *queryResolver) Movie(ctx context.Context, id string) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MoviesByIds(ctx context.Context, ids []string) ([]*model.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context, page *int, title *string, genre *string, director *string, producer *string) (*model.Movies, error) {
	pageLimit := 20
	var count int64

	var moviesGraph []*model.Movie
	var movies []models.Movie
	// Query to get the count
	database.DB.Find(&movies, "title LIKE ? AND genre LIKE ? AND director LIKE ? AND producer LIKE ?", "%"+*title+"%", "%"+*genre+"%", "%"+*director+"%", "%"+*producer+"%").Count(&count)
	// Query to get the results
	database.DB.Limit(pageLimit).Offset((*page-1)*pageLimit).Preload("Actors").Find(&movies, "title LIKE ? AND genre LIKE ? AND director LIKE ? AND producer LIKE ?", "%"+*title+"%", "%"+*genre+"%", "%"+*director+"%", "%"+*producer+"%")
	for _, movie := range movies {
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
		movie := movieAssembler(movie)
		movie.Actors = actorsGraph
		moviesGraph = append(moviesGraph, &movie)
	}
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
	return &model.Movies{
		Info: &model.Info{
			Count:    &countPointer,
			Last:     &last,
			Next:     &next,
			Previous: &previous,
		},
		Data: moviesGraph,
	}, nil
}

func (r *queryResolver) Actor(ctx context.Context, id string) (*model.Actor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ActorsByIds(ctx context.Context, ids []string) ([]*model.Actor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Actors(ctx context.Context, page *int, name *string, gender *string) (*model.Actors, error) {
	var actorsGraph []*model.Actor
	var actors []models.Actor
	database.DB.Preload("Movies").Find(&actors)
	for _, actor := range actors {
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
		actor := actorAssembler(actor)
		actor.Movies = moviesGraph
		actorsGraph = append(actorsGraph, &actor)
	}
	return &model.Actors{
		Data: actorsGraph,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
