package web

import (
	"html/template"
	"net/http"
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

func getServices() *[]Service {
	services := []Service{}
	services = append(services, Service{
		Name:   "main branch",
		Url:    "https://github.com/the-developer-guy/LoTo",
		Locked: false,
	})
	services = append(services, Service{
		Name:   "coffee machine",
		Url:    "https://www.cl.cam.ac.uk/coffee/xvcoffee.html",
		Locked: true,
	})

	return &services
}

func StartServer() {
	tmpl := template.Must(template.ParseFiles("web/templates/main.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := MainPage{
			PageTitle: "LoTo - 🔒Lockout-Tagout🏷️",
			Version:   "1.0.0",
			Services:  getServices(),
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
