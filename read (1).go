package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

/*
 * Your program should prompt the user for the name of the text file.
 * Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
 * Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
 * After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
 */
func main() {

	names := make([]Name, 0)

	var fileName string
	fmt.Println("Enter the file name: ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		fileName = scanner.Text()
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstAnsLastName := strings.Fields(line)
		name := Name{Fname: firstAnsLastName[0], Lname: firstAnsLastName[1]}
		names = append(names, name)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, value := range names {
		fmt.Printf("First name: %s, Last name: %s \n", value.Fname, value.Lname)
	}

}
