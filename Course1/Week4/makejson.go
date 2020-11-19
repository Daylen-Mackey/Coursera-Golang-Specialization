package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//Write a program
//which prompts the
//user to first enter a
//name, and then enter an
//address. Your program should
//create a map and add the name and
//address to the map using the keys “name”
//and “address”, respectively. Your program should
//use Marshal() to create a JSON object from the map, and
//then your program should print the JSON object.



func main() {

	//Prompting name input

	fmt.Println("Please enter a name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputName := scanner.Text()

	//Prompting address input
	fmt.Println("Please enter an address:")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputAddress := scanner.Text()

	//declaring our map
	inputMap := make(map[string]string)

	inputMap["name"] = inputName
	inputMap["address"] = inputAddress


	fmt.Println("The original map: ",inputMap)

	jsonObject,_ := json.Marshal(inputMap)

	fmt.Println("The JSON object: ",jsonObject)
	fmt.Println("The JSON object as a string: ",string(jsonObject))

	var decodedMap map[string]string
	err := json.Unmarshal(jsonObject, &decodedMap)

	fmt.Println("Unmarshaled JSON: ",decodedMap," Error: ",err)
}