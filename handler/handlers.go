package handler

import (
	"net/http"

	"github.com/anas-dev-92/AGA/config"
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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//here we can build some logic functionality
	stringmap := make(map[string]string)
	stringmap["hello"] = "from the string map"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringmap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringmap,
	}) // here send the data to the template
}
func (m *Repository) Resources(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "Resources.page.tmpl", &models.TemplateData{})
}
func (m *Repository) UploadData(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "UploadData.page.tmpl", &models.TemplateData{})
}
