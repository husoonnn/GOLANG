package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func Swap2 (numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}


func BubbleSort2 (numbers []int) {
	// fmt.Print(len(numbers))
	sorted := true

	for sorted {
		sorted = false
		for i := 0; i < len(numbers)-1; i++ {
			// fmt.Print(i)
			if numbers[i] > numbers[i+1] {
				Swap2(numbers, i)
				sorted = true
			}
		}
	}
	
}


func main2() {
	var numbers []int
	fmt.Print("Enter some integers: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	for _, s := range strings.Fields(scanner.Text()) {
		i, err := strconv.Atoi(s)
		// Only accept first 10 integers 
		if err == nil && len(numbers) < 10 {
			numbers = append(numbers, i)
		}
	}
	BubbleSort2(numbers)
	fmt.Println(numbers)

}
