package auth

import (
	"go1/db"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)
func HalamanRegister(w http.ResponseWriter, r *http.Request){
	Tmpl.ExecuteTemplate(w, "register.html", nil)
}
func Register(w http.ResponseWriter, r *http.Request){
	var a Akun
	a.NamaDepan  = r.FormValue("namaDepan")
	a.NamaBelakang = r.FormValue("namaBelakang")
	a.Email = r.FormValue("email")
	a.Username = r.FormValue("username")
	a.Password = r.FormValue("password")

	// AFTER:
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
    	http.Error(w, "Gagal memproses password", http.StatusInternalServerError)
     return
	}
	_, err = db.DB.Exec("Insert into users (nama_depan, nama_belakang, username, email, password, role) value (?, ?, ?, ?, ?, ?)",
    	a.NamaDepan, a.NamaBelakang, a.Username, a.Email, string(hash), "users")
	if err != nil {
    	http.Error(w, "Gagal menyimpan akun", http.StatusInternalServerError)
     return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
