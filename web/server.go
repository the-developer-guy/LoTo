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
	http.HandleFunc("/lock", Lock)
	http.HandleFunc("/unlock", Unlock)

	var err error
	db, err = internals.NewDatabaseHelper()
	if err != nil {
		return
	}

	err = http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}

func Home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, homePage)
}

func Lock(w http.ResponseWriter, r *http.Request) {
	db.Lock("coffee machine")
	w.Write([]byte("ok"))
}

func Unlock(w http.ResponseWriter, r *http.Request) {
	db.Unlock("coffee machine")
	w.Write([]byte("ok"))
}
