// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Actor struct {
	ID     *string  `json:"id"`
	Name   *string  `json:"name"`
	Age    *int     `json:"age"`
	Gender *string  `json:"gender"`
	Movies []*Movie `json:"movies"`
}

type Actors struct {
	Info *Info    `json:"info"`
	Data []*Actor `json:"data"`
}

type FilterActor struct {
	Name *string `json:"name"`
}

type FilterMovie struct {
	Title    *string `json:"title"`
	Genre    *string `json:"genre"`
	Director *string `json:"director"`
	Producer *string `json:"producer"`
}

type Info struct {
	Count    *int `json:"count"`
	Pages    *int `json:"pages"`
	Next     *int `json:"next"`
	Previous *int `json:"previous"`
}

type Movie struct {
	ID       *string  `json:"id"`
	Title    *string  `json:"title"`
	Genre    *string  `json:"genre"`
	Price    *string  `json:"price"`
	Director *string  `json:"director"`
	Producer *string  `json:"producer"`
	Actors   []*Actor `json:"actors"`
}

type Movies struct {
	Info *Info    `json:"info"`
	Data []*Movie `json:"data"`
}
