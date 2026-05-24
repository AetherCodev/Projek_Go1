package users

import (
	"fmt"
	"go1/auth"
	"go1/db"
	"html/template"
	"net/http"
)

var Tmpl *template.Template


type Produk struct {
	IdProduk int
	NamaProduk string
	Harga      int
	Stok       int
}

func DashboardUser(w http.ResponseWriter, r *http.Request) {
	koneksi := db.DB
	session, _ := auth.Store.Get(r, "login-session")
	role, ok := session.Values["role"].(string)
	if !ok ||role != "users" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	rows, err := koneksi.Query("Select id_produk, nama_produk, harga, stok from produk")
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
		err := rows.Scan(&p.IdProduk, &p.NamaProduk, &p.Harga, &p.Stok)
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

	err = Tmpl.ExecuteTemplate(w, "users/index.html", map[string]interface{}{
    	"Produk": daftarProduk,
	})
	if err != nil {
    	fmt.Println("Error render template:", err.Error())
     http.Error(w, "Gagal render halaman", http.StatusInternalServerError)
     return
	}
}
