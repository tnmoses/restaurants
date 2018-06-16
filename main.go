package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/asdine/storm"
	reflections "gopkg.in/oleiade/reflections.v1"
)

func main() {
	http.HandleFunc("/restaurants", restaurants)
	http.HandleFunc("/restaurants/", restaurant)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", nil)
}

func healtcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All is well")
}

func restaurant(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/restaurants/")
	id, err := strconv.Atoi(path)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	switch r.Method {
	case "GET":
		getone(id, w)
	case "PUT":
		update(id, w, r)
	case "DELETE":
		delete(id, w)
	default:
		respondWithError(w, http.StatusBadRequest, "Unsupported HTTP method")
	}
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		list(w)
	case "POST":
		create(w, r)
	default:
		respondWithError(w, http.StatusBadRequest, "Unsupported HTTP method")
	}
}

func update(id int, w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant
	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// get restaurant that should be updated
	err = db.One("ID", id, &restaurant)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Restaurant not found")
		return
	}

	// parse the JSON
	jsonMap := make(map[string]interface{})
	if err = json.NewDecoder(r.Body).Decode(&jsonMap); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// update restaurant struct with data contained in jsonMap
	for key, value := range jsonMap {
		if key != "ID" {
			err = reflections.SetField(&restaurant, key, value)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}
		} else {
			respondWithError(w, http.StatusBadRequest, "Cannot update primary key")
			return
		}
	}

	// save update restaurant struct
	err = db.Save(&restaurant)
	if err != nil {
		respondWithError(w, http.StatusNotModified, "Restaurant not update")
		return
	}
	respondWithJSON(w, http.StatusOK, restaurant)
}

func delete(id int, w http.ResponseWriter) {
	var restaurant Restaurant
	restaurant.ID = id

	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.DeleteStruct(&restaurant)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Restaurant not found")
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}

func getone(id int, w http.ResponseWriter) {
	var restaurant Restaurant
	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.One("ID", id, &restaurant)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Restaurant not found")
		return
	}
	respondWithJSON(w, http.StatusOK, restaurant)
}

func list(w http.ResponseWriter) {
	var restaurants []Restaurant
	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.AllByIndex("ID", &restaurants)
	if err != nil {
		panic(err)
	}
	respondWithJSON(w, http.StatusOK, restaurants)
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
