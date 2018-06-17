package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/asdine/storm"
	reflections "gopkg.in/oleiade/reflections.v1"
)

// Update handles http PUT requests to a single resource
func Update(id int, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s", "UPDATE", r.RequestURI, "Update")
	var restaurant Restaurant
	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// get restaurant that should be updated
	err = db.One("ID", id, &restaurant)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Restaurant not found")
		return
	}

	// parse the JSON
	jsonMap := make(map[string]interface{})
	if err = json.NewDecoder(r.Body).Decode(&jsonMap); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// update restaurant struct with data contained in jsonMap
	for key, value := range jsonMap {
		if key != "ID" {
			err = reflections.SetField(&restaurant, key, value)
			if err != nil {
				RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}
		} else {
			RespondWithError(w, http.StatusBadRequest, "Cannot update primary key")
			return
		}
	}

	// save update restaurant struct
	err = db.Save(&restaurant)
	if err != nil {
		RespondWithError(w, http.StatusNotModified, "Restaurant not updated")
		return
	}
	RespondWithJSON(w, http.StatusOK, restaurant)
}

func Delete(id int, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s", "DELETE", r.RequestURI, "Delete")
	var restaurant Restaurant
	restaurant.ID = id

	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.DeleteStruct(&restaurant)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Restaurant not found")
		return
	}
	RespondWithJSON(w, http.StatusNoContent, nil)
}

func GetOne(id int, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s", "GET", r.RequestURI, "GetOne")
	var restaurant Restaurant
	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.One("ID", id, &restaurant)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Restaurant not found")
		return
	}
	RespondWithJSON(w, http.StatusOK, restaurant)
}

func List(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s", "GET", r.RequestURI, "List")
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
	RespondWithJSON(w, http.StatusOK, restaurants)
}

func Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s", "POST", r.RequestURI, "Create")
	var restaurant Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
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
	RespondWithJSON(w, http.StatusCreated, restaurant)
}
