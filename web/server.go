package web

import (
	"html/template"
	"net/http"
)

type MainPage struct {
	PageTitle string
	Version   string
}

func StartServer() {
	tmpl := template.Must(template.ParseFiles("web/templates/main.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := MainPage{
			PageTitle: "LoTo - 🔒Lockout-Tagout🏷️",
			Version:   "1.0.0",
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
