package handlers

import (
	"net/http"
	"web_01/pkg/config"
	"web_01/pkg/models"
	"web_01/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request)  {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{} )
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request)  {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}


