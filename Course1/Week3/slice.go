package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

//
//Write a program which prompts the user to enter integers and stores
//the integers in a sorted slice. The program should be written as a
//loop. Before entering the loop, the program should create an empty
//integer slice of size (length) 3. During each pass through the loop,
//the program prompts the user to enter an integer to be added to the
//slice. The program adds the integer to the slice, sorts the slice,
//and prints the contents of the slice in sorted order. The slice must
//grow in size to accommodate any number of integers which the user
//decides to enter. The program should only quit (exiting the loop)
//when the user enters the character ‘X’ instead of an integer.

func main() {
	// Creating the slice of length three
	programSlice := make([]int,0,3)
	for{
		fmt.Println("Enter an integer, or type 'X' (uppercase) to end ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		if userInput == "X"{
			fmt.Println("Your Final Slice:")
			fmt.Println(programSlice)
			break
		}else{
			var intInput, err = strconv.Atoi(userInput)
			if err == nil{
				programSlice = append(programSlice,intInput )
				sort.Ints(programSlice)
				fmt.Println(programSlice)
			}else{
				fmt.Println("Not a valid entry. Please try again")

			}

		}


	}


	//var inputFloat float64
	//fmt.Println("Please enter a float:")
	//fmt.Scan(&inputFloat)

}