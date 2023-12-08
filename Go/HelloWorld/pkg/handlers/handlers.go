package handlers

import (
	"net/http"

	"github.com/snhkn/100DaysOfCode/Go/HelloWorld/pkg/config"
	"github.com/snhkn/100DaysOfCode/Go/HelloWorld/pkg/render"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
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
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &TemplateData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	//perform some business logic
	//get a data
	strMap := make(map[string]string)
	strMap["test"] = "Hello, again"

	//send the data to the template

	render.RenderTemplate(w, "about.page.tmpl", &TemplateData{
		StringMap: strMap,
	})
}
