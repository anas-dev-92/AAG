package main

import (
	"log"

	"github.com/anas-dev-92/AGA/config"

	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/anas-dev-92/AGA/handler"
	"github.com/anas-dev-92/AGA/render"
)

const PORT = ":8080"

var app config.AppConfig        // first create the variable
var session *scs.SessionManager //to use the session in all the app
func main() {
	app.Inproduction = false
	session = scs.New()
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Inproduction
	session.Lifetime = 24 * time.Hour

	app.Session = session

	tc, err := render.CreateTemplateCache() // then create the template cache
	if err != nil {
		log.Fatal("can't create template")
	}
	app.TemplateCache = tc //assign the template cache inside the app.templatecache
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	render.NewTemplate(&app) // last is to call the function and put the templatecache inside of it

	ser := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	err = ser.ListenAndServe()
	log.Fatal(err)
}
