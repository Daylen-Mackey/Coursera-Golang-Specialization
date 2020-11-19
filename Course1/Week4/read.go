package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//
//Write a program which reads information
//from a file and represents it in a slice
//of structs. Assume that there is a text
//file which contains a series of names.
//
//	Each line of the text file has a
//first name and a last name, in that order,
//separated by a single space on the line.
//
//Your program will define a name struct
//which has two fields, fname for the first
//name, and lname for the last name. Each
//field will be a string of size 20 (characters).
//
//Your program should prompt the user for
//the name of the text file. Your program
//will successively read each line of the
//text file and create a struct which contains
//the first and last names found in the file.
//Each struct created will be added to a
//slice, and after all lines have been read
//from the file, your program will have a
//slice containing one struct for each line
//in the file. After reading all lines from the
//file, your program should iterate through your
//slice of structs and print the first and last
//names found in each struct.

func main() {
	//My defined name file is names.txt
	type FullName struct {
		fname string
		lname string
	}

	nameSlice := make([]FullName,0,10) //Give initial cap of 10
	fmt.Println("Please enter your file ('filename.txt'):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()

	file,_ := os.Open(fileName)
	scanner = bufio.NewScanner(file)
	for scanner.Scan(){
		splitString := strings.Split(scanner.Text()," ")
		var tempName FullName
		tempName.fname = splitString[0]
		tempName.lname = splitString[1]
		nameSlice = append(nameSlice,tempName)
	}
	file.Close()

	//Time to iterate through the slice

	for _, name := range nameSlice{

		fmt.Println(name.fname, name.lname)
	}













}