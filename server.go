package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func LoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func main() {
	app := mux.NewRouter()
	app.HandleFunc("/", HomePage)
	app.HandleFunc("/form", LoginPage)

	log.Fatal(http.ListenAndServe(":8080", app))
	fmt.Printf("Server listening at localhost:8080")
}
