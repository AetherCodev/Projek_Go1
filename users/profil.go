	package users

import (
	"go1/auth"
	"go1/db"
	"net/http"
	"fmt"
)

func ProfilUsers(w http.ResponseWriter, r *http.Request){
	koneksi := db.DB
	session, _ := auth.Store.Get(r, "login-session")
	id, okId:= session.Values["id_users"].(int)
	role, ok := session.Values["role"].(string)
	if !okId {
    	http.Redirect(w, r, "/login", http.StatusSeeOther)
    	return
	}
	if !ok ||role != "users"{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	row := koneksi.QueryRow(
    "SELECT nama_depan, nama_belakang, username, email FROM users WHERE id_users = ?", id,
)
	var a auth.Akun
err := row.Scan(&a.NamaDepan, &a.NamaBelakang, &a.Username, &a.Email)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
execErr := Tmpl.ExecuteTemplate(w, "users/profil.html", a)
if execErr != nil {
    fmt.Println("Error render profil:", execErr.Error())
    http.Error(w, "Gagal render halaman", http.StatusInternalServerError)
}
}
