package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const token = "Z3JhbWEgbW9saGFkYSB0ZW0gY2hlaXJvIGJvbQ=="

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", Index)
	router.HandleFunc("/auth", AuthHandler)
	router.HandleFunc("/login", LoginHandler)

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	log.Printf("Server is running on http://localhost:5000/")
	log.Fatal(server.ListenAndServe())
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nada aqui, mas você pode acessar /login e /auth, tem coisa lá =)")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	type Response struct {
		Token string `json:"token"`
	}

	response := Response{token}

	json.NewEncoder(w).Encode(response)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	type Response struct {
		Logged bool `json:"logged"`
	}

	online := Response{true}
	offline := Response{false}

	auth := strings.Fields(r.Header.Get("authorization"))

	var userToken string

	if len(auth) >= 2 {
		userToken = auth[1]
	}

	if userToken == token {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(online)
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(offline)
	}
}
