package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/mecamon/booking/pkg/config"
	"github.com/mecamon/booking/pkg/models"
	"github.com/mecamon/booking/pkg/render"
	"log"
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

	render.RenderTemplate(writer, req, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(writer http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "This is a test data for about page"

	remoteIp := m.AppConfig.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(writer, req, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Generals(writer http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(writer, req, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(writer http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(writer, req, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(writer http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(writer, req, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(writer http.ResponseWriter, req *http.Request) {
	start := req.Form.Get("start")
	end := req.Form.Get("end")
	writer.Write([]byte(fmt.Sprintf("The posted values are: %s and %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJson(writer http.ResponseWriter, req *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")

	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(out)
}

func (m *Repository) MakeReservation(writer http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(writer, req, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(writer http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(writer, req, "contact.page.tmpl", &models.TemplateData{})
}
