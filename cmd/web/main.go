package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NotAliAhmad/learningGO/pkg/config"
	"github.com/NotAliAhmad/learningGO/pkg/handlers"
	"github.com/NotAliAhmad/learningGO/pkg/render"

)

const portNumber = ":8080"


// main function 
func main(){
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("Cannot create template cache",err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.Newhandlers(repo)
	render.NewTemplates(&app)
	//first arg is the path and second is the hander func
	http.HandleFunc("/",handlers.Repo.Homepage)
	http.HandleFunc("/about",handlers.Repo.Aboutpage)

	// control + c to to stop the service
	log.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	// Starts the service
	_ = http.ListenAndServe(portNumber,nil) 
}
