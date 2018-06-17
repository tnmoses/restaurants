package main

type Restaurant struct {
	ID          int `storm:"id,increment"` // primary key
	Name        string
	Phone       string
	Cuisines    string // csv
	Address     string
	Description string
}
