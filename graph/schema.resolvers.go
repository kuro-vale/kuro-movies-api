package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/graph/generated"
	"github.com/kuro-vale/kuro-movies-api/graph/model"
	"github.com/kuro-vale/kuro-movies-api/models"
)

func (r *queryResolver) Movie(ctx context.Context, id string) (*model.Movie, error) {
	var movieGraph model.Movie
	var movie models.Movie
	database.DB.Preload("Actors").Find(&movie, "id = ?", id)
	actorsGraph := nestedActors(movie)
	movieGraph = movieAssembler(movie)
	movieGraph.Actors = actorsGraph
	return &movieGraph, nil
}

func (r *queryResolver) MoviesByIds(ctx context.Context, ids []string) ([]*model.Movie, error) {
	var moviesGraph []*model.Movie
	for _, id := range ids {
		var movie models.Movie
		database.DB.Preload("Actors").Find(&movie, "id = ?", id)
		actorsGraph := nestedActors(movie)
		movieGraph := movieAssembler(movie)
		movieGraph.Actors = actorsGraph
		moviesGraph = append(moviesGraph, &movieGraph)
	}
	return moviesGraph, nil
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
		actorsGraph := nestedActors(movie)
		movie := movieAssembler(movie)
		movie.Actors = actorsGraph
		moviesGraph = append(moviesGraph, &movie)
	}
	info := generateInfo(page, count, pageLimit)
	return &model.Movies{
		Info: &info,
		Data: moviesGraph,
	}, nil
}

func (r *queryResolver) Actor(ctx context.Context, id string) (*model.Actor, error) {
	var actorGraph model.Actor
	var actor models.Actor
	database.DB.Preload("Movies").Find(&actor, "id = ?", id)
	moviesGraph := nestedMovies(actor)
	actorGraph = actorAssembler(actor)
	actorGraph.Movies = moviesGraph
	return &actorGraph, nil
}

func (r *queryResolver) ActorsByIds(ctx context.Context, ids []string) ([]*model.Actor, error) {
	var actorsGraph []*model.Actor
	for _, id := range ids {
		var actor models.Actor
		database.DB.Preload("Movies").Find(&actor, "id = ?", id)
		moviesGraph := nestedMovies(actor)
		actorGraph := actorAssembler(actor)
		actorGraph.Movies = moviesGraph
		actorsGraph = append(actorsGraph, &actorGraph)
	}
	return actorsGraph, nil
}

func (r *queryResolver) Actors(ctx context.Context, page *int, name *string, gender *string) (*model.Actors, error) {
	pageLimit := 20
	var count int64

	var actorsGraph []*model.Actor
	var actors []models.Actor
	// Query to get the count
	database.DB.Find(&actors, "name LIKE ? AND gender ILIKE ?", "%"+*name+"%", *gender).Count(&count)
	// Query to get the results
	database.DB.Limit(pageLimit).Offset((*page-1)*pageLimit).Preload("Movies").Find(&actors, "name LIKE ? AND gender ILIKE ?", "%"+*name+"%", *gender)
	for _, actor := range actors {
		moviesGraph := nestedMovies(actor)
		actor := actorAssembler(actor)
		actor.Movies = moviesGraph
		actorsGraph = append(actorsGraph, &actor)
	}
	info := generateInfo(page, count, pageLimit)
	return &model.Actors{
		Info: &info,
		Data: actorsGraph,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
