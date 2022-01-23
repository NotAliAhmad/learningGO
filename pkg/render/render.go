package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/NotAliAhmad/learningGO/pkg/config"
)

var functions = template.FuncMap{
}

var app *config.AppConfig

// sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}

// Renders templates using html templates
func RenderTemplate(w http.ResponseWriter, tmpl string){
// The template cache should come from the app config 
	var tc map[string]*template.Template 
	if app.UseCache{
		tc = app.TemplateCache
	}else{
		tc, _ = CreateTemplateCache()
	}


	// error check if an invalid entry is put then it should stop the program
	t, wrongtmpl := tc[tmpl] 
	if !wrongtmpl {
		log.Fatal("Could not get template from template cache")
	}

	// creating a new bytes buffer and executing the template
	buf := new(bytes.Buffer)
	_ = t.Execute(buf,nil)

	// take the buffer and write it to the response writer
	// response writer is the page itself and we are writing the html to the page
	_,err := buf.WriteTo(w)
	if err != nil{
		fmt.Println("Error writing template to browser", err)
	}
 }


// Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template,error){

	// creating a map of type string key and template pointer value
	myCache := map[string]*template.Template{}

	// Glob returns all the files given in the string param
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil{
		return myCache, err
	}

	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return myCache,err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil{
			return myCache,err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil{
				return myCache,err
			}
		}
		myCache[name] = ts
	}
	return myCache,nil
}