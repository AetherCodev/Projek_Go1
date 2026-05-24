package users

import (
	"go1/auth"
	"go1/db"
	"net/http"
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
	row, err := koneksi.Query("Select nama_depan, nama_belakang, username, email from users where id_users = ?", id)
	if err != nil { // ← pakai err di sini! ✅
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
	}
	defer row.Close()
	var Profil []auth.Akun
	for row.Next() {
		var a auth.Akun
		err := row.Scan(&a.NamaDepan, &a.NamaBelakang, &a.Username, &a.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return	
		}
		Profil = append(Profil, a)
	}
	err = Tmpl.ExecuteTemplate(w, "users/profil.html", map[string]interface{}{
    	"Profil": Profil,
	})
	if err != nil {
     	http.Error(w, "Gagal render halaman", http.StatusInternalServerError)
      	return
	}
}
