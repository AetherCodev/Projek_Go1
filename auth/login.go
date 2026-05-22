package auth

import (
	"go1/db"
	"html/template"
	"net/http"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)
type Akun struct {
	idUsers int
	namaDepan string
	namaBelakang string
	Username string
	Email string
	Password string
	Role string
}
var Tmpl *template.Template
var Store =sessions.NewCookieStore([]byte("anggajoki1"))
func DashboardAdmin(w http.ResponseWriter, r *http.Request) {
    Tmpl.ExecuteTemplate(w, "admin/index.html", nil)
}

func DashboardUser(w http.ResponseWriter, r *http.Request) {
    Tmpl.ExecuteTemplate(w, "user/index.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	koneksi := db.DB
	username := r.FormValue("username")
	password := r.FormValue("password")
	rows, err := koneksi.Query("select id_users,username, password, role from users where username = ?", username)
	if err != nil { // ← pakai err di sini! ✅
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
	}
	defer rows.Close()

	var cekAkun []Akun
	for rows.Next() {
		var a Akun
		err := rows.Scan(&a.idUsers, &a.Username, &a.Password, &a.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cekAkun = append(cekAkun, a)
	}
	if len(cekAkun) == 0 {
    // Kirim pesan error ke template
    	Tmpl.ExecuteTemplate(w, "login.html", map[string]string{
        	"Error": "Username or password is'nt valid",
    	})
    	return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(cekAkun[0].Password),
		[]byte(password),
	)

	if err != nil {
		Tmpl.ExecuteTemplate(w, "login.html", map[string]string{
        "Error": "Username or password is'nt valid",
    })
    return 
	}
	
	session, _ := Store.Get(r, "login-session")
	session.Values["username"] = cekAkun[0].Username
	session.Values["id_users"] = cekAkun[0].idUsers
	session.Values["role"] = cekAkun[0].Role
	session.Save(r, w)
	
	if cekAkun[0].Role == "admin"{
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/users/dashboard", http.StatusSeeOther)
	}
}

func HalamanLogin(w http.ResponseWriter, r *http.Request) {
    // Tampilkan login.html
    Tmpl.ExecuteTemplate(w, "login.html", nil)
}
