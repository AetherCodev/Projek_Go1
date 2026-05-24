package db
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func Koneksi(){
	var err error
	// AFTER:
	DB, err = sql.Open("mysql", "root:@AnGgA123@tcp(localhost:3306)/web2")
	if err != nil {
    	fmt.Println("Gagal parse DSN:", err)
     return
	}
	if err = DB.Ping(); err != nil {
    	fmt.Println("Gagal konek ke database:", err)
     return
	}
	fmt.Println("Koneksi berhasil")
}