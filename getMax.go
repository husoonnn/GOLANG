// Variable argument number
// functions can take a variable number of arguments
// use ellipsis ... to specify
// Treated as a slice inside function
package main
import "fmt"

func getMax(vals ...int) int{
	maxV:=-1
	for _,v := range (vals){
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main(){
	defer fmt.Println("max found!") // defer is only printed after the whole main block is done
	//however if some form of variable is defined and is used, it will print whatever is operated
	// on the variable in sequence
	fmt.Print(getMax(100,2,3), "\n")
	vslice := []int{100,2,3} //can use ellipsis on array as well
	fmt.Print(getMax(vslice...), "\n")
}