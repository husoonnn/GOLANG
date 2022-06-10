package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct{
	Name string `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct{
	EntryNo int `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct{
	Name string `json:"name"`
}

func main(){
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println(response)
	responseData, err := ioutil.ReadAll(response.Body) //returns a stream of bytes
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("ResponseObject", responseObject)
	fmt.Println("Name of responseObject", responseObject.Name)
	fmt.Println("Total Number of Pokemons:", len(responseObject.Pokemon))
	fmt.Println(responseObject.Pokemon)


	fmt.Println("All pokemons: ")
	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(i, responseObject.Pokemon[i].Species.Name)
	}
}
