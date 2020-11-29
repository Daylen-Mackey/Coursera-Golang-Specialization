package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

*/

// Philosopher will think, get hungry, eat, then repeat.
// The number of philosophers and the number of servings they're allowed can be modified

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	number int
	leftChopstick, rightChopstick *Chopstick
	servingsEaten int
}

func (self Philosopher) pickupChopsticks(){
	if rand.Intn(2) == 0{

		self.leftChopstick.Lock()
		self.rightChopstick.Lock()

	}else{
		self.rightChopstick.Lock()
		self.leftChopstick.Lock()

	}

}
func (self Philosopher) putdownChopsticks(){
	self.rightChopstick.Unlock()
	self.leftChopstick.Unlock()

}



func (self Philosopher) eat(numOfServings int,wg *sync.WaitGroup) {
	for {
		self.pickupChopsticks()

		fmt.Println("Philosopher",self.number,"started eating" )
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		self.servingsEaten++

		self.putdownChopsticks()

		fmt.Println("Philosopher",self.number,"finished eating" )
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

		if self.servingsEaten == numOfServings{
			fmt.Println("Philosopher", self.number, "is done eating. Servings:", self.servingsEaten)
			break
		}
	}
	wg.Done()
}


var diningWaitGroup sync.WaitGroup

func main() {
	// How many philosophers and servings we're using
	numberOfPhilosophers := 5
	numberOfServingsAllowedInt := 3

	// Create chopsticks
	chopstickSlice := make([]*Chopstick, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		chopstickSlice[i] = new(Chopstick)
	}

	// Creating the philosophers and assigning them chopsticks
	philosophers := make([]*Philosopher, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		philosophers[i] = &Philosopher{i, chopstickSlice[i],chopstickSlice[(i+1)%numberOfPhilosophers],0}
		diningWaitGroup.Add(1)
		go philosophers[i].eat(numberOfServingsAllowedInt,&diningWaitGroup)

	}

	diningWaitGroup.Wait() // Wait for all the goroutines to finish

}