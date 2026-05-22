package db
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func Koneksi(){
	var err error
	DB, err = sql.Open("mysql", "root:@AnGgA123@tcp(localhost:3306)/web2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Koneksi berhasil")
}