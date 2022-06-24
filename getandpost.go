package main

import (
	"log"
	"net/http"
)

type book struct {
	title         string
	author        author
	publiation    string
	publishedDate string
}
type author struct {
	firstname string
	lastname  string
	dob       string
	penname   string
}

func GetOne(w http.ResponseWriter, r *http.Request) {

}

func GetAll(w http.ResponseWriter, r *http.Request) {

}
func PostBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.HandleFunc("/book/1", GetOne)
	log.Fatal(http.ListenAndServe(":8000", nil))
	http.HandleFunc("/book", GetAll)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
