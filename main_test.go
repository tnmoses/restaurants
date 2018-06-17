package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testRestaurant = []byte(`
{
	"name": "Mitte Kebab",
	"Phone": "+4917112345623",
	"Cuisines": "turkish, spicy",
	"Address": "Mitte",
	"Description": "late night food"
}
`)

func TestCreateHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/restaurants", bytes.NewBuffer(testRestaurant))
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restaurants)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Test the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Decode response into expected Restaurant type
	var restaurant Restaurant
	err = json.NewDecoder(rr.Body).Decode(&restaurant)
	if err != nil {
		t.Fatal(err)
	}
	if restaurant.ID == 0 ||
		restaurant.Phone == "" ||
		restaurant.Cuisines == "" ||
		restaurant.Address == "" ||
		restaurant.Description == "" {
		t.Errorf("handler returned unexpected body: missing fields")
	}
}

func TestListHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/restaurants", nil)
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restaurants)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Test the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode response into expected map of Restaurants
	var restaurants []Restaurant
	err = json.NewDecoder(rr.Body).Decode(&restaurants)
	if err != nil {
		t.Fatal(err)
	}

	// Map should not be empty because the create test runs above
	// fmt.Printf("TestListHandler: %v\n", restaurants)
	if len(restaurants) == 0 {
		t.Errorf("handler returned empty array")
	}
}

// test get One
func TestGetOneandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/restaurants/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restaurant)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Test the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// fmt.Printf("Get one handler: %s\n", rr.Body.String())

	// Decode response into expected Restaurant type
	var restaurant Restaurant
	err = json.NewDecoder(rr.Body).Decode(&restaurant)
	if err != nil {
		t.Fatal(err)
	}

	if restaurant.ID == 0 ||
		restaurant.Phone == "" ||
		restaurant.Cuisines == "" ||
		restaurant.Address == "" ||
		restaurant.Description == "" {
		t.Errorf("handler returned unexpected body: missing fields")
	}
}

// test PUT
func TestPutHandler(t *testing.T) {
	var putQuery = []byte(`{"Description": "actually, bad food"}`)
	req, err := http.NewRequest("PUT", "/restaurants/1", bytes.NewBuffer(putQuery))
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restaurant)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Test the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode response into expected Restaurant type
	var restaurant Restaurant
	err = json.NewDecoder(rr.Body).Decode(&restaurant)
	if err != nil {
		t.Fatal(err)
	}

	if restaurant.Description != "actually, bad food" {
		t.Errorf("incorrect description field. expected %s, got %s", string(putQuery), restaurant.Description)
	}
}

//test DELETE
func TestDeleteHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/restaurants/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(restaurant)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Test the status code
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

	// get all restaurants to ensure the first was deleted
	req, err = http.NewRequest("GET", "/restaurants", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(restaurants)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var restaurants []Restaurant
	err = json.NewDecoder(rr.Body).Decode(&restaurants)
	if err != nil {
		t.Fatal(err)
	}

	// Map should not be empty because the create test runs above
	// fmt.Printf("TestListHandler: %v\n", restaurants)
	if restaurants[0].ID == 1 {
		t.Errorf("first item was not deleted")
	}

}
