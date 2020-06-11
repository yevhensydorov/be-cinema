package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// GetMovies function is for getting movies array from api
func GetMovies(w http.ResponseWriter, r *http.Request) {
	movieAPIURL := os.Getenv("API_URL")
	movieAPIKey := os.Getenv("API_KEY")

	vars := mux.Vars(r)
	movieName := vars["movieName"]

	fmt.Println(movieAPIURL + movieAPIKey + "&s=" + movieName)
}
