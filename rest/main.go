package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
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

func handleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/articles", ArticlesPage)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func main() {

	Articles = []Article{
		Article{
			Title:  "Title 1 to example",
			Author: "Nosirov Otabek",
			Link:   "obdev.com",
		},
		Article{
			Title:  "Title 2 to example",
			Author: "Mister X",
			Link:   "obdev.com",
		},
	}

	handleRequests()
}
