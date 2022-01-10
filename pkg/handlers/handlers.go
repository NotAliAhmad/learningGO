package handlers

import (
	"net/http"

	"github.com/NotAliAhmad/learningGO/pkg/config"
	"github.com/NotAliAhmad/learningGO/pkg/render"
)
// Repository variable
var Repo *Repository

// This is the repository type
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig)*Repository{
	return &Repository{
		App : a,
	}
}
// sets the repository for the handlers
func Newhandlers(r * Repository){
	Repo = r
}

// This is the homepage handler
// All handlers must accept 2 params a reponsewriter and http request pointer
func (m *Repository) Homepage(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.html")
}

// This is the about handler
func (m *Repository) Aboutpage(w http.ResponseWriter, r *http.Request){
   render.RenderTemplate(w, "about.page.html")
}