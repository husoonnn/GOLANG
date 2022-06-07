/*
Write a program which reads information from a file and represents it in a slice of structs. Assume 
that there is a text file which contains a series of names. Each line of the text file has a first 
name and a last name, in that order, separated by a single space on the line. 

Your program will define a name struct which has two fields, fname for the first name, and lname 
for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively 
read each line of the text file and create a struct which contains the first and last names found in 
the file. Each struct created will be added to a slice, and after all lines have been read from the 
file, your program will have a slice containing one struct for each line in the file. After reading 
all lines from the file, your program should iterate through your slice of structs and print the
first and last names found in each struct.
*/
package main
import(
	"fmt"
	//"ioutil"
	"os"
	"bufio"
	"strings"
)
type Person struct{
	fname string
	lname string
}
func main(){
	slice := make([]Person, 0, 2)
	var input string
	fmt.Println("Input file name: ")
	fmt.Scan(&input)
	f, err := os.Open(input)
	if err != nil {
        fmt.Println(err)
    }
	defer f.Close()
    // read the file word by word using scanner
    scanner := bufio.NewScanner(f)
    //scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        // do something with a word
        s := strings.Split(scanner.Text(), " ")
		var person Person
		person.fname, person.lname = s[0], s[1]
		slice = append(slice, person)
    }
	if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
	//to access each slice accordinglyfmt.Println(slice[2].lname)
	for _, v := range slice {
		fmt.Printf("First name: %s, Last name: %s \n", v.fname, v.lname)
	}
}