package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Nome     string    `json:"name"`
	Pokemons []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	numero  int            `json:"entry_number"`
	especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Nome string `json:"name"`
}

func main() {
	res, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/") // response mapeado en JSON
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Nome)
	fmt.Println(len(responseObject.Pokemons))

	for _, Pokemon := range responseObject.Pokemons {
		fmt.Println(Pokemon.especie.Nome)
	}
}
