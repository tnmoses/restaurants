package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asdine/storm"
)

func main() {
	http.HandleFunc("/restaurants", restaurants)
	http.HandleFunc("/restaurants/", getone)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", nil)
}

func healtcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All is well")
}

func getone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get one")
	//check if is int, if not fallback to restaurants
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "get list of restaurants")
	case "POST":
		create(w, r)
	case "PUT":
		fmt.Fprintln(w, "update existing restaurant")
	case "DELETE":
		fmt.Fprintln(w, "delete existing restaurant")
	default:
		fmt.Fprintln(w, "unsupported http method")
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Save(&restaurant)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", restaurant)
	respondWithJSON(w, http.StatusCreated, restaurant)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type Restaurant struct {
	ID          int `storm:"id,increment"` // primary key
	Name        string
	Phone       string
	Cuisines    string // csv
	Address     string
	Description string
}
