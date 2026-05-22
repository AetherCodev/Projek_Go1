package main

import (
	"fmt"
	"html/template"
	"net/http"
	"go1/db"
	"github.com/gorilla/mux"
	"go1/auth"
	"go1/users"
)

var tmpl = template.Must(
    template.New("").ParseGlob("templates/*.html"),
)

func init() {
    template.Must(tmpl.ParseGlob("templates/users/*.html"))
    template.Must(tmpl.ParseGlob("templates/admin/*.html"))
}

func main() {
	db.Koneksi()
	auth.Tmpl = tmpl
	users.Tmpl = tmpl
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
    r.PathPrefix("/static/").Handler(
        http.StripPrefix("/static/", fs),
    )
    r.HandleFunc("/login", auth.HalamanLogin).Methods("GET")
    r.HandleFunc("/login", auth.Login).Methods("POST")
    r.HandleFunc("/register", auth.HalamanRegister).Methods("GET")
    r.HandleFunc("/register", auth.Register).Methods("POST")
    r.HandleFunc("/users/dashboard",users.DashboardUser).Methods("GET")
    fmt.Println("Server jalan di port 8080")
    http.ListenAndServe(":8080", r)
}