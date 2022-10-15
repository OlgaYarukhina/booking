package handlers

import (
	"net/http"

	"github.com/OlgaYarukhina/booking/pkg/config"
	"github.com/OlgaYarukhina/booking/pkg/models"
	"github.com/OlgaYarukhina/booking/pkg/render"
)

//Repo is the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppCongig
}
//NewRepo creates a new repository 
func NewRepo (a *config.AppCongig) *Repository {
	return &Repository{
		App: a,
	}
}
//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// або tmpl, _ := template.ParseFiles(("templates/home.page.html"))
	render.RenderTemlate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	//m.App.Session.
	stringMap["remote_ip"] = remoteIP

	//send data to the tamplate
	render.RenderTemlate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

