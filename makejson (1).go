package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
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

	jsonFromMap, _ := json.Marshal(personMap)

	fmt.Println(string(jsonFromMap))

}
