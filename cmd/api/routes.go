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
	mux.Use(app.enableCORS)       //enable CORS

	// endpoints

	mux.Get("/homepage", app.HomePage)
	mux.Get("/status", app.Statuss)
	mux.Get("/all-movies", app.Movies)

	mux.Post("/authenticate", app.authenticate)

	//return teh mux
	return mux
}
