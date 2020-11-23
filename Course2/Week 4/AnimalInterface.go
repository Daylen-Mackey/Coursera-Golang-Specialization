package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

/*
Write a program which allows the user to create a set of animals and to get
information about those animals. Each animal has a name and can be either a
cow, bird, or snake. With each command, the user can either create a new
animal of one of the three types, or the user can request information about
an animal that he/she has already created. Each animal has a unique name,
defined by the user. Note that the user can define animals of a chosen type,
but the types of animals are restricted to either cow, bird, or snake. The
following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, “>”, to indicate that the
user can type a request. Your program should accept one command at a time from
the user, print out a response, and print out a new prompt on a new line. Your
program should continue in this loop forever. Every command from the user must
be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first
string is “newanimal”. The second string is an arbitrary string which will be the
name of the new animal. The third string is the type of the new animal, either “cow”,
“bird”, or “snake”. Your program should process each newanimal command by creating the new
animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal. The third string is the name of the information
requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query
command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the
Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and
return no values. The Eat() method should print the animal’s food, the Move() method should print the
animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types
Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that
the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create
an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
*/

/*
Approach
---------
1) User allowed to create animals
	- Can only be cow, bird, or snake though
		- That's where assertion comes into play
2) Leading input is a command
	- Either 'newanimal' or ''query'
3) We've got an interface called "Animal", then 3 types of animals



---------------------------------------------------------------------------------
Note: I would never take this approach if given a task with similar requirements
The requirements force you to write additional unnecessary code, but here we are
---------------------------------------------------------------------------------


*/
// Animal Interface
type Animal interface {
	eat()
	move()
	speak()
}
// Cow Object
type Cow struct{
	name string
	food string
	movement string
	sound string
}

func (self *Cow) initCow(name string){
	self.name = name
	self.food = "Grass"
	self.movement = "Walk"
	self.sound = "Moooooooo"
}

func (self Cow) Eat(){
	fmt.Println(self.food)
}
func (self Cow) Move(){
	fmt.Println(self.movement)
}
func (self Cow) Speak(){
	fmt.Println(self.sound)
}

// Bird Object
type Bird struct{
	name string
	food string
	movement string
	sound string
}

func (self *Bird) initBird(name string){
	self.name = name
	self.food = "Worms"
	self.movement = "Fly"
	self.sound = "Peep"
}

func (self Bird) Eat(){
	fmt.Println(self.food)
}
func (self Bird) Move(){
	fmt.Println(self.movement)
}
func (self Bird) Speak(){
	fmt.Println(self.sound)
}


// Snake Object
type Snake struct{
	name string
	food string
	movement string
	sound string
}

func (self Snake) Eat(){
	fmt.Println(self.food)
}
func (self Snake) Move(){
	fmt.Println(self.movement)
}
func (self Snake) Speak(){
	fmt.Println(self.sound)
}
func (self *Snake) initSnake(name string){
	self.name = name
	self.food = "Mice"
	self.movement = "Slither"
	self.sound = "Hsss"
}

func main() {

	animalMap := make(map[string]interface{})


	// Get Input:
	fmt.Println("Hello, you have two options. Please make sure to use a single space for separation")
	fmt.Println("1) Create a new animal: 'newanimal <name> <animal type>")
	fmt.Println("2) Query an animal you created: 'query <animal name> <eat, move, or speak>")

	acceptedMethods := []string{"Move", "Speak", "Eat"}

	for{
		fmt.Print("\n>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputStringSlice := strings.Split(scanner.Text()," ")
		var userCommand string  = strings.ToLower(inputStringSlice[0])
		var animalName string = strings.ToLower(inputStringSlice[1])
		if userCommand == "newanimal"{
			var animalType string = strings.ToLower(inputStringSlice[2])

			switch animalType{
				case "cow":
					var cowInstance Cow
					cowInstance.initCow(animalName)
					animalMap[animalName]  = cowInstance
					fmt.Println("Created It!")

				case "bird":
					var birdInstance Bird
					birdInstance.initBird(animalName)
					animalMap[animalName]  = birdInstance
					fmt.Println("Created It!")

				case "snake":
					var snakeInstance Snake
					snakeInstance.initSnake(animalName)
					animalMap[animalName]  = snakeInstance
					fmt.Println("Created It!")

			default:
				fmt.Println("Not a valid animal type.")
			}

		}else if (userCommand == "query"){
			var inputMethod string = strings.Title(inputStringSlice[2])

			_, found := Find(acceptedMethods, inputMethod)
			if found{
				method2Call := reflect.ValueOf(animalMap[animalName]).MethodByName(inputMethod)
				method2Call.Call(nil)
			}else{
				fmt.Println("Please use an accepted Method", acceptedMethods)
			}

		}else{
			fmt.Println("Please use a valid command (newanimal or query)")
		}

	}






}

// Function from https://golangcode.com/check-if-element-exists-in-slice/
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}