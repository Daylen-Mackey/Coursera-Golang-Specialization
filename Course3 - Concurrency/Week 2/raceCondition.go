package main

import "fmt"


//Write two goroutines which have a race condition when executed concurrently.
//Explain what the race condition is and how it can occur.

/*
Explanation:
- A race condition is when results, outcome, or behaviour of a program/syste, is dependent on timing and/or sequencing

- In the following program, I run four incrementers (each having a maximum of 100000).
Three of these incrementers are run with goroutines, while the last one is run normally.

The incrementer function directly alters the input variable, so the final statement prints the sum of the current
state of all four incrementer inputs. The race condition occurs because different goroutines may be in different stages
of incrementing the value. If this were just run sequentially, you'd get a output total of 100000 * 4 = 400000. Because
three incrementers are run with goroutines, different interleavings will give different outputs nearly every time.


*/


func incrementer(x *int, numOfIncrements int){
	var valuePointer *int
	valuePointer = x
	for i := 0; i < numOfIncrements; i++ {
		*valuePointer++
		//fmt.Println(*valuePointer)

	}


}

func main() {

	var firstNum int = 0
	var secondNum int = 0
	var thirdNum int = 0
	var fourthNum int = 0
	go incrementer(&firstNum, 100000)
	go incrementer(&secondNum, 100000)
	go incrementer(&thirdNum, 100000)
	incrementer(&fourthNum, 100000)

	finalVal := firstNum + secondNum + thirdNum + fourthNum
	fmt.Println(finalVal)




}
