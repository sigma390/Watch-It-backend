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

	mux.Get("/homepage", app.HomePage)
	mux.Get("/status", app.Statuss)

	//return teh mux
	return mux
}
