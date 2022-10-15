package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"github.com/OlgaYarukhina/booking/pkg/config"
	"github.com/OlgaYarukhina/booking/pkg/models"
)


var app *config.AppCongig

//NewTemplates sets config for hte template package
func NewTemplates(a *config.AppCongig){
app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate render a template
func RenderTemlate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	
var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the app config
	tc = app.TemplateCache
	}else{
		tc, _ = CreateTemplateCache()
	}
	
	          //get request template fron cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)


	
	             //render the template
    _, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from /templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	//range through allfiles ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}


//create a template cache


//render the template


/*
var tc = make(map[string]*template.Template)

func RenderTemlate (w http.ResponseWriter, t string) {
var tmpl *template.Template
var err error

_, inMap := tc[t]
if !inMap {
log.Println("Create template and addingto cache")
//need to create the template
err = createTemplateCache(t)
if err != nil {
	log.Println(err)
}

}else{
	log.Println("Using cached template")
}

tmpl = tc[t]

err = tmpl.Execute(w, nil)
if err != nil {
	log.Println(err)
}
}

func createTemplateCache(t string) error {
	templates := [] string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}
	// parse the template
	//var templates [] string
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)

	tc[t] = tmpl

	return nil
}
*/