package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	http.HandleFunc("/", enableCors(HomeHandler))
	http.HandleFunc("/welcome", enableCors(WelcomeHandler))
	fmt.Println("[*] Server running at port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the home page of this API!"))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API!"))
}

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}
