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

func runOperation(i int, j int, k int, allNumbersChosen []int) int {
	operations := []string{"+", "-", "*", "/"}
	numberChosen := allNumbersChosen[i]
	nextNumberChosen := allNumbersChosen[j]
	operationChosen := operations[k]

	// Call `doCalculation` with chosen numbers and operation
	return doCalculation(numberChosen, nextNumberChosen, operationChosen)
}

func doCalculation(numberChosen int, nextNumberChosen int, operation string) int {
	switch operation {
	case "+":
		return numberChosen + nextNumberChosen
	case "-":
		return numberChosen - nextNumberChosen
	case "/":
		if nextNumberChosen != 0 {
			return numberChosen / nextNumberChosen
		}
		return 0
	case "*":
		return numberChosen * nextNumberChosen
	default:
		return 0
	}
}

func generateCombinations(allNumbersChosen []int, size int) [][]int {
	var combinations [][]int
	var helper func(int, []int)

	helper = func(start int, combo []int) {
		if len(combo) == size {
			// Make a copy of combo and add it to combinations
			comboCopy := make([]int, len(combo))
			copy(comboCopy, combo)
			combinations = append(combinations, comboCopy)
			return
		}
		for i := start; i < len(allNumbersChosen); i++ {
			helper(i+1, append(combo, allNumbersChosen[i]))
		}
	}

	helper(0, []int{})
	return combinations
}

func applyOperations(numbers []int, currentResult int, index int, operations []string) {
	if index == len(numbers) {
		// We've reached the end, print the result
		fmt.Printf("Result of expression: %d\n", currentResult)
		return
	}

	// Apply each operation to the next number in the sequence
	for _, op := range operations {
		newResult := doCalculation(currentResult, numbers[index], op)
		fmt.Printf("Applying %d %s %d = %d\n", currentResult, op, numbers[index], newResult)
		applyOperations(numbers, newResult, index+1, operations)
	}
}

func calculateOnCombinations(allNumbersChosen []int) {
	operations := []string{"+", "-", "*", "/"}
	for size := 2; size <= len(allNumbersChosen); size++ {
		combinations := generateCombinations(allNumbersChosen, size)
		for _, combo := range combinations {
			fmt.Printf("Calculating for combination: %v\n", combo)
			applyOperations(combo, combo[0], 1, operations)
		}
	}
}

func main() {
	fmt.Println("Welcome to countdown")

	var numberOfBigNumber int
	fmt.Print("Enter the number of big numbers: ")
	fmt.Scan(&numberOfBigNumber)

	// Seed the random generator once at the beginning of the program
	rand.Seed(time.Now().UnixNano())

	// Assume `chooseBigNumbers` and `chooseSmallNumbers` are defined functions
	// Get the chosen big numbers
	allNumbersChosen := chooseBigNumbers(numberOfBigNumber)
	allNumbersChosen = chooseSmallNumbers(allNumbersChosen)

	fmt.Println("Numbers chosen are:", allNumbersChosen)

	// Perform calculations on all combinations
	calculateOnCombinations(allNumbersChosen)
}
