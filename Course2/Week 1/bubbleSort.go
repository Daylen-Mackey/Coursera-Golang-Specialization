package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Description:
//------------
/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest. Use your
favorite search tool to find a description of how the bubble sort algorithm works.
As part of this program, you should write a function called BubbleSort() which takes a slice of integers as
an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.
A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two
adjacent elements in the slice. You should write a Swap() function which performs this operation. Your
Swap() function should take two arguments, a slice of integers and an index value i which indicates a
position in the slice. The Swap() function should return nothing, but it should swap the contents of
the slice in position i with the contents in position i+1.
Submit your Go program source code.
*/

/*
Steps:
1) User types sequence up to 10 integers
2) Print the input
3) Sort
	- Use BubbleSort() function
		- Takes a slice of integers and returns nothing
	- Uses Swap() function
		- Takes slice of integers, and index value as arguments
		- swap i with i + 1
		- returns nothing
4) Print the sorted output

*/

func stringSlice2IntSlice(stringSlice []string) []int {
	sliceLength := len(stringSlice)
	intSlice := make([]int,0,sliceLength)
	for _, val := range stringSlice {
		intVal, _ := strconv.Atoi(val)
		intSlice = append(intSlice, intVal)
	}
	return intSlice
}

func Swap(slice []int, index int){
	smaller := slice[index + 1]
	larger := slice[index]
	slice[index] = smaller
	slice[index + 1] = larger

}


func BubbleSort(slice []int){
	sliceLength := len(slice)
	for{ //for is equivalent to while loop
		var swapsMade int = 0
		for i := 0; i < sliceLength - 1; i++ {
			if slice[i] > slice[i + 1]{
				Swap(slice,i)
				swapsMade++
				//fmt.Println(slice)
			}

		}

		if swapsMade == 0{
			break
		}
	}



}



func main() {
	fmt.Println("Please enter your sequence of integers ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Get the input, convert it to an int slice, print it for the user
	inputStringSlice := strings.Split(scanner.Text()," ")
	inputIntSlice := stringSlice2IntSlice(inputStringSlice)
	fmt.Println("Your input:",inputIntSlice)
	BubbleSort(inputIntSlice)
	fmt.Println("Sorted Output:",inputIntSlice)



}
