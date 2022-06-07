/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/
package main

import (
	"fmt"
	
)

func BubbleSort(sli []int) {
	for i := 0; i < len(sli); i++{
		for j := 0; j < len(sli) - 1 - i; j++{
			if sli[j + 1] < sli[j] {
				Swap(sli, j)
			}	
		}
	}	
}

func Swap(numbers []int, j int){
	temp := numbers[j]
	numbers[j] = numbers[j+1]
	numbers[j+1] = temp
}

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func main(){
	fmt.Println("Enter input:")
	arr := input([]int{}, nil)
	fmt.Println("Input:", arr)
	BubbleSort(arr)
	fmt.Println(arr)
	
}
