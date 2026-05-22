package auth

import (
	"fmt"
	"go1/db"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)
func HalamanRegister(w http.ResponseWriter, r *http.Request){
	Tmpl.ExecuteTemplate(w, "register.html", nil)
}
func Register(w http.ResponseWriter, r *http.Request){
	var a Akun
	a.namaDepan  = r.FormValue("namaDepan")
	a.namaBelakang = r.FormValue("namaBelakang")
	a.Email = r.FormValue("email")
	a.Username = r.FormValue("username")
	a.Password = r.FormValue("password")

	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	_, err = db.DB.Exec("Insert into users (nama_depan, nama_belakang, username, email, password, role) value (?, ?, ?, ?, ?, ?)",
	a.namaDepan, a.namaBelakang, a.Username, a.Email, string(hash), "users")
	if err != nil {
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
