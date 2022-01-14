package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/mecamon/booking/pkg/config"
	"github.com/mecamon/booking/pkg/handlers"
	"github.com/mecamon/booking/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	appConfig.Session = session

	templateCache, err := render.CreateTemplateCached()

	if err != nil {
		log.Fatal("Error caching templates: ", err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false

	//Passing the reference to pointer in the other file
	render.NewTemplate(&appConfig)
	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}

	fmt.Println("Server started in port:", portNumber)

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start the server: ", err)
	}

}
