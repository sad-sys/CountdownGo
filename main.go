package main

import (
	"fmt"
	"math/rand"
	"time"
)

func chooseBigNumbers(numberOfBigNumber int) []int {
	bigNumbers := []int{25, 50, 75, 100}
	numbersChosen := []int{}

	for i := 0; i < numberOfBigNumber; i++ {
		randomIndex := rand.Intn(len(bigNumbers)) // Use the globally seeded rand.Intn
		randomElement := bigNumbers[randomIndex]
		numbersChosen = append(numbersChosen, randomElement)
		bigNumbers = append(bigNumbers[:randomIndex], bigNumbers[randomIndex+1:]...) // Remove chosen element
	}

	return numbersChosen
}

func chooseSmallNumbers(allNumbersChosen []int) []int {
	// Add small numbers to make a total of 6 numbers
	for len(allNumbersChosen) < 6 {
		randomNumber := rand.Intn(10) + 1 // Random number between 1 and 10
		allNumbersChosen = append(allNumbersChosen, randomNumber)
	}
	return allNumbersChosen
}

func main() {
	fmt.Println("Welcome to countdown")

	var numberOfBigNumber int
	fmt.Print("Enter the number of big numbers: ")
	fmt.Scan(&numberOfBigNumber)

	// Seed the random generator once at the beginning of the program
	rand.Seed(time.Now().UnixNano())

	// Get the chosen big numbers
	allNumbersChosen := chooseBigNumbers(numberOfBigNumber)

	allNumbersChosen = chooseSmallNumbers(allNumbersChosen)

	fmt.Println("Number is ", allNumbersChosen)
}
