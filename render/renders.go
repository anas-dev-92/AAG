package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/anas-dev-92/AGA/config"
	"github.com/anas-dev-92/AGA/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{} //It is used to define a mapping of function names to functions that can be used within Go templates.
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func DefultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// the where is what server will response back to the client
func RenderTemplate(w http.ResponseWriter, r *http.Request, temp string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache { // this is when ask if the app is up to date then take it from cache else create new copy
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[temp]
	if !ok {
		log.Fatal("could not get the template from templateCache")
	}
	buf := new(bytes.Buffer)
	td = DefultData(td, r)
	_ = t.Execute(buf, td) // here we exec the templates and we store the data inside the buffer bytes
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}

// this function will render all the pages in the templates folder and combine it with the layout
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCach := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl") // here the return (pages) is type of slice or map so have index and value
	if err != nil {
		return myCach, err
	}
	for _, page := range pages {
		name := filepath.Base(page) // here we just notify the page that we will used it in the handlers
		fmt.Println("you're currently on the page :", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page) // ts is stands for template set
		if err != nil {
			return myCach, err
		}

		//here we will check if there is any layout find, so we can match it with our templates set inside the ts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCach, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCach, err
			}
		}
		myCach[name] = ts
	}
	return myCach, nil
}
