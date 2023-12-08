package handlers

import (
	"github.com/yenmaximoff/bookings/pkg/config"
	"github.com/yenmaximoff/bookings/pkg/models"
	"github.com/yenmaximoff/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	StringMap := make(map[string]string)
	StringMap["test"] = "Hello, everybody"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}
