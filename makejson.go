/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name”
and “address”, respectively. Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)
type Person struct{
	Name string `json:"Name"`
	Address string `json:"Address"`
}

func main(){
	/*p1 := map[string]interface{}{
		name: "joe", 
		address: "Latha. St",
	}*/
	var personMap = make(map[string]string)
	var input string
	fmt.Println("Write a name: ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}
	personMap["name"] = input

	fmt.Println("Write an address: ")

	if scanner.Scan() {
		input = scanner.Text()
	}

	personMap["address"] = input
	
	p1, err := json.Marshal(personMap)
	if err != nil { 
		fmt.Println(err)
	}
	fmt.Println(string(p1))
}