package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/yenmaximoff/bookings/pkg/config"
	"github.com/yenmaximoff/bookings/pkg/handlers"
	"github.com/yenmaximoff/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//измените это на true во время производства
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", port)
	//_ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
