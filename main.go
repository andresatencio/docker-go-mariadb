package main
import (
	"net/http"
	"io"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}


func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":5000", nil)
}