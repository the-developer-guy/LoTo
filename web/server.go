package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/the-developer-guy/LoTo/internals"
)

type MainPage struct {
	PageTitle string
	Version   string
	Services  *[]Service
}

type Service struct {
	Name   string
	Url    string
	Locked bool
}

func StartServer(config *internals.LotoConfig) {
	services := []Service{}
	for _, service := range *config.Services {
		s := Service{
			Name:   service.Name,
			Url:    service.Url,
			Locked: false,
		}
		services = append(services, s)
	}
	tmpl := template.Must(template.ParseFiles("web/templates/main.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := MainPage{
			PageTitle: config.Name,
			Version:   "1.0.0",
			Services:  &services,
		}
		tmpl.Execute(w, data)
	})

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}
