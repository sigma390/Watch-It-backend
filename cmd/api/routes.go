package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	//mux
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer) //defualt middleware To log panicks

	// endpoints

	//return teh mux
	return mux
}
