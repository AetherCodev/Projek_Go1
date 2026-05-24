package users

import (
	"fmt"
	"go1/db"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var Tmpl *template.Template
var Store = sessions.NewCookieStore([]byte("anggajoki1"))

type Produk struct {
	NamaProduk string
	Harga      int
	Stok       int
}

func DashboardUser(w http.ResponseWriter, r *http.Request) {
	koneksi := db.DB
	session, _ := Store.Get(r, "login-session")
	role := session.Values["role"]
	if role != "users" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	rows, err := koneksi.Query("Select nama_produk, harga, stok from produk")
	if err != nil {
		fmt.Println("Error query:", err.Error())
		Tmpl.ExecuteTemplate(w, "users/index.html", map[string]string{
			"Error": err.Error(),
		})
		return
	}
	defer rows.Close()
	var daftarProduk []Produk
	for rows.Next() {
		var p Produk
		err := rows.Scan(&p.NamaProduk, &p.Harga, &p.Stok)
		if err != nil {
			fmt.Println("Error scan:", err.Error())
			Tmpl.ExecuteTemplate(w, "users/index.html", map[string]string{
				"Error": err.Error(),
			})
			return
		}
		daftarProduk = append(daftarProduk, p)
	}

	fmt.Println("Jumlah produk:", len(daftarProduk))
	if len(daftarProduk) == 0 {
		fmt.Println("Tabel produk kosong!")
	}

	Tmpl.ExecuteTemplate(w, "users/index.html", map[string]interface{}{
		"Produk": daftarProduk,
	})
}
