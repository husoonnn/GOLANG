package main

import (
	"fmt"
	"strconv"
)

func BubbleSort1(numbers []int){
	for i := 0; i < len(numbers); i++{
		for j := 0; j < len(numbers) - 1 - i; j++{
			if numbers[j + 1] < numbers[j] {
				Swap1(numbers, j)
			}	
		}
	}	
}
func Swap1(numbers []int, j int){
	temp := numbers[j]
	numbers[j] = numbers[j + 1]
	numbers[j + 1] = temp
}

func main1(){
	var s string
	var numbers []int
	fmt.Println("Please input 10 numbers:")	
	for i := 0; i < 10; i++ {
		fmt.Scan(&s)
		
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Input " + s + " is not a number! Discarding...")
			continue
		}
		numbers = append(numbers, n)
	}
	BubbleSort1(numbers)
	fmt.Println(numbers)
}

