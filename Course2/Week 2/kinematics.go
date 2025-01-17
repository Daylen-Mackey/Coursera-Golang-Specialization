package main

import (
	"fmt"
)

/*
Let us assume the following formula for displacement s as a function of time t,
acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration,
initial velocity, and initial displacement. Then the program should prompt the
user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes
three float64 arguments, acceleration a, initial velocity vo, and initial displacement
so. GenDisplaceFn() should return a function which computes displacement as a function
of time, assuming the given values acceleration, initial velocity, and initial displacement.

The function returned by GenDisplaceFn() should take one float64 argument t, representing time,
and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration,
initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to generate a function fn which will
compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)
Then I can use the following statement to print the displacement after 3 seconds.
fmt.Println(fn(3))
And I can use the following statement to print the displacement after 5 seconds.
fmt.Println(fn(5))
*/

func GenDisplaceFn(a float64, vo float64, so float64) func (float64) float64 {
	returnFunction := func (t float64) float64{
		//s =½ a t2 + vot + so

		return (.5 * a)*(t*t) + vo*(t) + so

	}

	return returnFunction
}


func main() {
	var acceleration float64
	var velocityInitial float64
	var displacementInitial float64
	var time float64

	fmt.Println("Please fill in the following values to solve the kinematics equation")
	fmt.Println("Acceleration:")
	fmt.Scan(&acceleration)
	fmt.Println("Initial Velocity:")
	fmt.Scan(&velocityInitial)
	fmt.Println("Initial Displacement:")
	fmt.Scan(&displacementInitial)
	fmt.Println("Time:")
	fmt.Scan(&time)

	displacementFunction := GenDisplaceFn(acceleration,velocityInitial,displacementInitial)
	finalDisplacement := displacementFunction(time)
	fmt.Println("Your Final Displacement:",finalDisplacement)

}