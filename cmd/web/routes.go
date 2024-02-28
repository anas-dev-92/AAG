package main

import (
	"net/http"

	"github.com/anas-dev-92/AGA/config"
	"github.com/anas-dev-92/AGA/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New() // here we make a new http handler which called mux and before we return this mux we start the routes for our app
	// mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	/* the diffrences between the chi and pat is that the chi include middlewares that will help us in work*/
	mux := chi.NewRouter()
	/*below here we put our middlewares.
	all meddilewares are found in this link https://github.com/go-chi/chi*/
	mux.Use(middleware.Recoverer) // this middleware are used to recover any panic that happen and print the reson behind it
	mux.Use(WriteToConsol)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/resources", handler.Repo.Resources)
	mux.Get("/UploadData", handler.Repo.UploadData)

	//this defenition to get static files for the pages
	Fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", Fileserver))
	return mux
}
