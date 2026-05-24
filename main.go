package main

import (
	"encoding/gob" 
	"fmt"
	"html/template"
	"net/http"
	"go1/db"
	"github.com/gorilla/mux"
	"go1/auth"
	"go1/users"
	"os"
)

var tmpl *template.Template

func init() {
    gob.Register(0)
    tmpl = template.Must(parseTemplates())
}

func parseTemplates() (*template.Template, error) {
    t := template.New("")
    files := []string{
        "templates/login.html",
        "templates/register.html",
    }
    subDirs := map[string]string{
        "templates/users/index.html": "users/index.html",
        "templates/admin/dashboard.html": "admin/dashboard.html",
    }
    // Parse root templates
    if _, err := t.ParseFiles(files...); err != nil {
        return nil, err
    }
    // Parse subfolder dengan nama eksplisit
    for path, name := range subDirs {
        b, err := os.ReadFile(path)
        if err != nil {
            return nil, err
        }
        if _, err = t.New(name).Parse(string(b)); err != nil {
            return nil, err
        }
    }
    return t, nil
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
    r.HandleFunc("/users/profil", users.ProfilUsers).Methods("GET")
    fmt.Println("Server jalan di port 8080")
    http.ListenAndServe(":8080", r)
}