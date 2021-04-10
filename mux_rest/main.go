package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Link   string `json:"Link"`
}

var Articles []Article

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func ArticlesPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: articlespage")
	json.NewEncoder(w).Encode(Articles)
}

func SingleArticlePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: articlespage with id")
	paths := mux.Vars(r)
	key := paths["id"]

	fmt.Println(key)

	for _, ar := range Articles {
		if ar.Id == key {
			json.NewEncoder(w).Encode(ar)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", ArticlesPage)
	myRouter.HandleFunc("/articles/{id}", SingleArticlePage)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {

	Articles = []Article{
		Article{
			Id:     "star",
			Title:  "Title 1 to example",
			Author: "Nosirov Otabek",
			Link:   "obdev.com",
		},
		Article{
			Id:     "moon",
			Title:  "Title 2 to example",
			Author: "Mister X",
			Link:   "obdev.com",
		},
	}

	handleRequests()
}
