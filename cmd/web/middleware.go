package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// this is the example of middleware that we gone hit whenever a page is clicked and you can find the implement of this in the "routes"
func WriteToConsol(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hitting the page")
		next.ServeHTTP(w, r)
	})

}

// add CSRF protection on all POST requests
func NoSurf(next http.Handler) http.Handler {
	crfHandler := nosurf.New(next)
	crfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.Inproduction,
		SameSite: http.SameSiteLaxMode,
	})
	return crfHandler
}

// Load and save session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
