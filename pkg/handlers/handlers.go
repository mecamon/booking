package handlers

import (
	"github.com/mecamon/booking/pkg/config"
	"github.com/mecamon/booking/pkg/models"
	"github.com/mecamon/booking/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	AppConfig *config.AppConfig
}

//NewRepo Creating the new repo
func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: appConfig,
	}
}

//NewHandlers sets the new repo
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(writer http.ResponseWriter, req *http.Request) {
	remoteIp := req.RemoteAddr
	m.AppConfig.Session.Put(req.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(writer http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "This is a test data for about page"

	remoteIp := m.AppConfig.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
