package users

import (
	// "go1/db"
    "html/template"
    "net/http"
    "github.com/gorilla/sessions"
)

var Tmpl *template.Template
var Store = sessions.NewCookieStore([]byte("anggajoki1"))

func DashboardUser(w http.ResponseWriter, r *http.Request) {

	session, _ := Store.Get(r, "login-session")
	role := session.Values["role"]
	if role != "users" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
    Tmpl.ExecuteTemplate(w, "users/index.html", nil)
}