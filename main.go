package main

import (
	"fmt"
	"math/rand"
	"time"
)

func smallNumbers(numberBig int) []int {
	rand.Seed(time.Now().UnixNano())

	smallNumbers := []int{}
	for i := 0; i < (5 - numberBig); i++ {
		randomSmall := rand.Intn(10) + 1 // Numbers between 1 and 10
		smallNumbers = append(smallNumbers, randomSmall)
	}
	return smallNumbers
}

func removeValue(s []int, value int) []int {
	for i, v := range s {
		if v == value {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s // Return original slice if value not found
}

func bigNumbers(numberBig int) []int {
	rand.Seed(time.Now().UnixNano())

	bigNumberToChoose := [4]int{25, 50, 75, 100}
	bigNumbersChosen := []int{}
	for i := 0; i < numberBig; i++ {
		bigNumberPosition := rand.Intn(4) // Random index between 0 and 3
		bigNumberChosen := bigNumberToChoose[bigNumberPosition]
		bigNumbersChosen = append(bigNumbersChosen, bigNumberChosen)
	}
	return bigNumbersChosen
}

func chooseNumbers(numberBig int) []int {
	smallNumbers := smallNumbers(numberBig)
	bigNumbers := bigNumbers(numberBig)

	allNumbers := append(bigNumbers, smallNumbers...)
	return allNumbers
}

func doOperations(startNumber int, allNumbers []int) int {
	rand.Seed(time.Now().UnixNano())
	operationNumber := rand.Intn(8) + 3 // Random number of operations between 3 and 10

	operations := []string{"-", "+", "/", "*"}
	for i := 0; i < operationNumber; i++ {
		if len(allNumbers) == 0 {
			break
		}
		randomIndex := rand.Intn(len(operations)) // Choose random operation
		randomOperation := operations[randomIndex]
		randomNumber := rand.Intn(len(allNumbers)) // Choose random number from list
		secondNumber := allNumbers[randomNumber]
		startNumber = doOperation(randomOperation, startNumber, secondNumber)
		allNumbers = removeValue(allNumbers, secondNumber)
	}

	return startNumber
}

func doOperation(randomOperation string, firstNumber int, secondNumber int) int {
	var result int
	switch randomOperation {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "/":
		if secondNumber != 0 {
			result = firstNumber / secondNumber
		} else {
			result = firstNumber // Handle division by zero by ignoring it
		}
	case "*":
		result = firstNumber * secondNumber
	default:
		fmt.Println("Unsupported operation")
		result = firstNumber
	}
	return result
}

func main() {
	var numberBig int

	fmt.Println("Enter the number of Big Numbers to Choose (0-4): ")
	fmt.Scan(&numberBig)

	if numberBig < 0 || numberBig > 4 {
		fmt.Println("Invalid input. Choose between 0 and 4.")
		return
	}

	finalResult := -1
	allNumbers := chooseNumbers(numberBig)

	for finalResult < 50 || finalResult > 999 {
		allNumbers := chooseNumbers(numberBig)

		startNumber := allNumbers[rand.Intn(len(allNumbers))]

		finalResult = doOperations(startNumber, allNumbers)
	}
	fmt.Println("Chosen Numbers:", allNumbers)
	fmt.Println("Final Result after Operations:", finalResult)
}
