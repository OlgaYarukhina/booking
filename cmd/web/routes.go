package main

import (
	"net/http"
	"github.com/OlgaYarukhina/booking/pkg/config"
	"github.com/OlgaYarukhina/booking/pkg/handlers"
	"github.com/go-chi/chi/v5" //https://github.com/go-chi/chi
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppCongig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/trip", handlers.Repo.Trip)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/bookingtrip", handlers.Repo.BookingTrip)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}

