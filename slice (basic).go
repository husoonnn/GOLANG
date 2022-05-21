/*Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an
empty integer slice of size (length) 3. During each pass through the loop, the program prompts the
user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts
the slice, and prints the contents of the slice in sorted order. The slice must grow in size to
accommodate any number of integers which the user decides to enter. The program should only quit
(exiting the loop) when the user enters the character ‘X’ instead of an integer. */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	mySlice := make([]int, 3)
	var input string
	
	for i := 0; i < 3; i++ {
		fmt.Println("Please enter an integer: ")
		fmt.Scanln(&input)
		if input == "X" || input == "x" {
			break
		}

		mySlice_var, err := strconv.Atoi(input)
		fmt.Print(mySlice_var, err)
		if err != nil { 
			fmt.Println(err)
			fmt.Println("Please input a valid integer!!")
			i--  //To repeat this loop
			continue
		}

		if i==len(mySlice) {
			mySlice = append(mySlice, mySlice_var)
		}else {
			mySlice[0] = mySlice_var
		}

		sort.Ints(mySlice)
		fmt.Println(mySlice)
	}
}