package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

/*
Write a program which allows the user to get information about a
predefined set of animals. Three animals are predefined, cow, bird,
and snake. Each animal can eat, move, and speak. The user can issue a
request to find out one of three things about an animal: 1) the food that
it eats, 2) its method of locomotion, and 3) the sound it makes when it
speaks. The following table contains the three animals and their associated
data which should be hard-coded into your program.


Animal	Food eaten	Locomotion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss
Your program should present the user with a prompt, “>”, to indicate that
the user can type a request. Your program accepts one request at a time from
the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever. Every request from the user
must be a single line containing 2 strings. The first string is the name of an
animal, either “cow”, “bird”, or “snake”. The second string is the name of the
information requested about the animal, either “eat”, “move”, or “speak”. Your
program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make
a type called Animal which is a struct containing three fields:food, locomotion,
and noise, all of which are strings. Make three methods called Eat(), Move(),
and Speak(). The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print
the animal’s locomotion, and the Speak() method should print the animal’s spoken
sound. Your program should call the appropriate method when the user makes a request.

Submit your Go program source code.


*/

/*
Notes:
--------------------
1) Three animals hard-coded
2) Present Prompt ">"
3) Loop forever
4) Single Line input, two strings


*/

type Animal struct{
	name string
	food string
	movement string
	sound string
}

func (self Animal) Eat(){
	fmt.Println(self.food)
}
func (self Animal) Move(){
	fmt.Println(self.movement)
}
func (self Animal) Speak(){
	fmt.Println(self.sound)
}


func main() {

	animalMap := make(map[string]Animal)
	animalMap["Cow"] = Animal{"Cow","Grass","Walk","Mooooooooooooooooooo"}
	animalMap["Bird"] = Animal{"Bird","Worms","Fly","Peep"}
	animalMap["Snake"] = Animal{"Snake","Mice","Slither","HSSSS"}

	// Get Input:
	fmt.Println("Please enter the animal and attribute separated by a space: 'Cow Eat' for example")
	fmt.Println("Manually Terminate when you are done")

	for{
		fmt.Print("\n>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputStringSlice := strings.Split(scanner.Text()," ")
		var inputAnimal string  = strings.Title(inputStringSlice[0])
		var inputMethod string = strings.Title(inputStringSlice[1])


		method2Call := reflect.ValueOf(animalMap[inputAnimal]).MethodByName(inputMethod)
		method2Call.Call(nil)
	}






}

