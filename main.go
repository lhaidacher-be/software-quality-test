package main

import (
	"fmt"
	"math/rand"
	"time"
)

var myArray [5]int

func main() {
	fmt.Println("Welcome to my program!")

	// Create a loop that populates the array with values
	for i := 0; i < 5; i++ {
		myArray[i] = rand.Intn(10)
	}

	// Print out the array
	fmt.Println("Here's the array:", myArray)

	// Wait for 2 seconds before continuing
	time.Sleep(time.Second * 2)

	// Call the processArray function
	sum := processArray(myArray)

	// Print out the result
	fmt.Println("The sum of the array is:", sum)
}

// Process the array and return the sum of its elements
func processArray(arr [5]int) int {
	sum := 0

	// Create a loop that calculates the sum of the array
	for _, val := range arr {
		sum = +val //FIXME
	}

	// Return the sum
	return sum
}
