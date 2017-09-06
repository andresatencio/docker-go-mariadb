package main
import (
	"net/http"
	"io"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}


func main() {
	fmt.Print("Hola busca 0.1")
	db, _ := sql.Open("mysql", "root:123456@/")
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":5000", nil)
}