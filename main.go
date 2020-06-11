package main

import (
	api "be-cinema/api"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies/{movieName}", api.GetMovies)
	http.Handle("/", r)

	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
