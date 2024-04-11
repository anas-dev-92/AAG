package handler

import (
	"log"
	"net/http"

	"github.com/anas-dev-92/AGA/config"
	froms "github.com/anas-dev-92/AGA/forms"
	"github.com/anas-dev-92/AGA/models"
	"github.com/anas-dev-92/AGA/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//here we can build some logic functionality
	stringmap := make(map[string]string)
	stringmap["hello"] = "from the string map"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringmap["remote_ip"] = remoteIp
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringmap,
	}) // here send the data to the template
}
func (m *Repository) Resources(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "Resources.page.tmpl", &models.TemplateData{})
}
func (m *Repository) UploadData(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "UploadData.page.tmpl", &models.TemplateData{
		Form: froms.New(nil), // now whenever we render this page we can get all the forms inside of it thus access any object in it
	})
}
func (m *Repository) PostUploadData(w http.ResponseWriter, r *http.Request) {
	// one of the first good practice to do is to sign an variable for error inisde r (http.request) to check
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	//its ok to start make all the variables inside the form here and work with it but better to create a new file for that which in our case model.go
	uploadDate := models.UploadData{
		FileName:       r.Form.Get("file_name"),
		PublishingDate: r.Form.Get("P_date"),
		Country:        r.Form.Get("country"),
		FileTypes:      r.Form.Get("file_type"),
	}
	form := froms.New(r.PostForm)

	form.Has("file_name", r)
	if !form.Valid() {
		data := make(map[string]interface{}) // we make this var so can hold the data that user will put in the form and its similer to data var in template file
		data["uploadData"] = uploadDate
		/*here we need to keep store the data if the user enter some of it true and valid so we can bring back the page with
		correct data he put and sent him error on the wrong one.
		here we call the same render for the uplaoddata page but we will not make new forms*/
		render.RenderTemplate(w, r, "UploadData.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return // this to stop processing this function
	}

}
func (m *Repository) Services(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "services.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
