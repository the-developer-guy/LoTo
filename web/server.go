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
	Services  *[]internals.Service
}

func StartServer(config *internals.LotoConfig) {
	tmpl := template.Must(template.ParseFiles("web/templates/main.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := MainPage{
			PageTitle: config.Name,
			Version:   "1.0.0",
			Services:  config.Services,
		}
		tmpl.Execute(w, data)
	})

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}
