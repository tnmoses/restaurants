package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/restaurants", restaurants)
	http.HandleFunc("/restaurants/", restaurant)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", nil)
}

func restaurant(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/restaurants/")
	id, err := strconv.Atoi(path)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	switch r.Method {
	case "GET":
		GetOne(id, w)
	case "PUT":
		Update(id, w, r)
	case "DELETE":
		Delete(id, w)
	default:
		RespondWithError(w, http.StatusBadRequest, "Unsupported HTTP method")
	}
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		List(w)
	case "POST":
		Create(w, r)
	default:
		RespondWithError(w, http.StatusBadRequest, "Unsupported HTTP method")
	}
}
