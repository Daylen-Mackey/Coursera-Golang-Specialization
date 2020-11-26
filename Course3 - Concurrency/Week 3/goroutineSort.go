package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

//sort.Ints()

/*
Write a program to sort an array of integers. The program should partition
the array into 4 parts, each of which is sorted by a different goroutine.

Each partition should be of approximately equal size. Then the main
goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each
goroutine which sorts ¼ of the array should print the subarray that it
will sort. When sorting is complete, the main goroutine should print the
entire sorted list.
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


func sortSlice(sliceInput []int, wg *sync.WaitGroup) {

	defer wg.Done()

	sort.Ints(sliceInput)


}

func main() {
	//Defining our wait group
	var wg sync.WaitGroup

	fmt.Println("Please enter your sequence of integers ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Get the input, convert it to an int slice, print it for the user
	inputStringSlice := strings.Split(scanner.Text()," ")
	inputIntSlice := stringSlice2IntSlice(inputStringSlice)
	fmt.Println("Your inputted array:",inputIntSlice)

	//For simplicity of this task
	sliceLength := len(inputIntSlice)
	separationPoint := sliceLength / 4.0


	// -- Following Code could be more elegant with looping, but this simple assignment is about wait groups
	//Only works for slices with length greater than 4
	group1 := inputIntSlice[0:separationPoint]
	group2 := inputIntSlice[separationPoint : separationPoint * 2]
	group3 := inputIntSlice[separationPoint * 2 : separationPoint * 3]
	group4 := inputIntSlice[separationPoint * 3 : sliceLength]


	fmt.Println("Grouped Inputs:",group1,group2,group3,group4)
	wg.Add(4)
	go sortSlice(group1, &wg)
	go sortSlice(group2, &wg)
	go sortSlice(group3, &wg)
	go sortSlice(group4, &wg)
	wg.Wait()
	fmt.Println("Sorted Groups:",group1,group2,group3,group4)
	var finalSlice []int
	finalSlice = append(finalSlice,group1...)
	finalSlice = append(finalSlice,group2...)
	finalSlice = append(finalSlice,group3...)
	finalSlice = append(finalSlice,group4...)
	fmt.Println("Appended Slice (Unsorted):", finalSlice)
	sort.Ints(finalSlice)
	fmt.Println("Final Sorted Slice:", finalSlice)

}