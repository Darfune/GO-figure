package main

import (
	"fmt"
	"log"
	"net/http"
	"web_01/pkg/config"
	"web_01/pkg/handlers"
	"web_01/pkg/render"
)

const portNumber = ":8080"

func main()  {
	var app config.AppConfig
	
	tc, err := render.CreateTemplateCashe()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
