package main

import (
	"fmt"
	"math/rand"
	"time"
)

func chooseBigNumbers( numberOfBigNumber int) vector[]

func main() {
	fmt.Println("Welcome to countdown")

	bigNumbers := []int{25, 50, 75, 100}
	numbersChosen := []int{}

	var numberOfBigNumber int
	fmt.Println("Enter the number of big numbers ")
	fmt.Scan(&numberOfBigNumber)

	for i := 0; i < numberOfBigNumber; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(source)
		randomIndex := randomGenerator.Intn(len(bigNumbers))
		// Choose the element at the random index
		randomElement := bigNumbers[randomIndex]
		numbersChosen = append(numbersChosen, randomElement)
		bigNumbers = append(bigNumbers[:randomIndex], bigNumbers[randomIndex+1:]...)
	}

}
