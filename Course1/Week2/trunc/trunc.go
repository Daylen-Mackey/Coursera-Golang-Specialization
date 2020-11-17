package main

/*
Write a program which prompts the user to enter a
floating point number and prints the integer
which is a truncated version of the floating
point number that was entered.

Truncation is the process of
removing the digits to the right of the decimal place.
*/
import "fmt"

func main() {
	var inputFloat float64
	fmt.Println("Please enter a float:")
	fmt.Scan(&inputFloat)

	var truncatedInput = int(inputFloat)
	fmt.Printf("Your Input: %f \nTruncated Output: %d", inputFloat, truncatedInput)

}
