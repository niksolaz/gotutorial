package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getUrlApiJSON(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		n := map[string]string{"name": responseObject.Pokemon[i].Species.Name, "url": responseObject.Pokemon[i].Species.Url}
		fmt.Fprintf(w, "%s\n", n)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome on Home Page!\n"))
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome on About Page!\n"))
	getUrlApiJSON(w, r)
}

func handleRequests() {
	port := ":8000"
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/about", AboutHandler)
	// In console see your basic path
	fmt.Println("http://localhost" + port)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	handleRequests()
}
