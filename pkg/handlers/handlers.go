package handlers

import (
	"net/http"

	"github.com/NotAliAhmad/learningGO/pkg/render"
)

// This is the homepage handler
// All handlers must accept 2 params a reponsewriter and http request pointer
func Homepage(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.tmpl")
}

// This is the about handler
func Aboutpage(w http.ResponseWriter, r *http.Request){
   render.RenderTemplate(w, "about.page.tmpl")
}