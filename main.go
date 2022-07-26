package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pet struct {
	Type         string `json:"type"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	FavoriteFood string `json:"favoritefood"`
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the pet matching server!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	pet := Pet{
		Type:         "Dog",
		Name:         "Cookie",
		Age:          2,
		FavoriteFood: "Pedigree",
	}

	if err := json.NewEncoder(w).Encode(pet); err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "Application/json")

}

func postHandler(w http.ResponseWriter, r *http.Request) {

	pet := Pet{}
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		fmt.Println(err)
	}

	fmt.Println(pet)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		fmt.Println(err)
	}

}

func badHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey pet lover! Unfortunately, the page you are looking for does not exist!")
}

func main() {
	fmt.Println("Server started ...")
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/register-pet", postHandler)
	http.HandleFunc("/get-pet", getHandler)
	http.HandleFunc("/", badHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
