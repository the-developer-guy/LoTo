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
	Services  *[]*internals.Service
}

var homeTemplate *template.Template
var homePage MainPage
var db *internals.DatabaseHelper

func StartServer(config *internals.LotoConfig) {
	homeTemplate = template.Must(template.ParseFiles("web/templates/main.html"))
	homePage = MainPage{
		PageTitle: config.Name,
		Version:   "1.0.0",
		Services:  config.Services,
	}
	http.HandleFunc("/", Home)
	http.HandleFunc("/lock/{name}", Lock)
	http.HandleFunc("/unlock/{name}", Unlock)

	var err error
	db, err = internals.NewDatabaseHelper()
	if err != nil {
		return
	}

	err = http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}

func Home(w http.ResponseWriter, r *http.Request) {
	var err error
	homePage.Services, err = db.GetServices()
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	homeTemplate.Execute(w, homePage)
}

func Lock(w http.ResponseWriter, r *http.Request) {
	serviceName := r.PathValue("name")
	db.Lock(serviceName)
	w.Header().Add("Content-Type", "")
	http.Redirect(w, r, fmt.Sprintf("http://%s", r.Host), http.StatusSeeOther)
}

func Unlock(w http.ResponseWriter, r *http.Request) {
	serviceName := r.PathValue("name")
	db.Unlock(serviceName)
	w.Header().Add("Content-Type", "")
	http.Redirect(w, r, fmt.Sprintf("http://%s", r.Host), http.StatusSeeOther)
}
