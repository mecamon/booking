package render

import (
	"bytes"
	"fmt"
	"github.com/mecamon/booking/pkg/config"
	"github.com/mecamon/booking/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var appConfig *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	appConfig = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

func RenderTemplate(writer http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	var templateCached map[string]*template.Template
	if appConfig.UseCache {
		templateCached = appConfig.TemplateCache
	} else {
		templateCached, _ = CreateTemplateCached()
	}

	t, ok := templateCached[tmpl]
	if !ok {
		log.Fatal("Error receiving template from appConfig")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)

	_ = t.Execute(buf, templateData)

	_, err := buf.WriteTo(writer)
	if err != nil {
		fmt.Println("Error printing template to browser:", err)
	}
}

var functions = template.FuncMap{}

func CreateTemplateCached() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			_, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
