package main

import (
	"fmt"
	"net/http"
)

func main() {
	// r := mux.NewRouter()
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
		fmt.Fprintln(w, "create new restaurant")
	case "PUT":
		fmt.Fprintln(w, "update existing restaurant")
	case "DELETE":
		fmt.Fprintln(w, "delete existing restaurant")
	default:
		fmt.Fprintln(w, "unsupported http method")
	}
}

// func restaurantDetail(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "find restaurants wip")
// }
//
// func createRestaurant(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "create restaurants wip")
// }
//
// func updateRestaurant(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "update restaurants wip")
// }
//
// func deleteRestaurant(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "delete restaurants wip")
// }
