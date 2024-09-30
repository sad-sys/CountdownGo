package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unicode"
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
		fmt.Println("+", firstNumber+secondNumber)
	case "-":
		result = firstNumber - secondNumber
		fmt.Println("-", firstNumber-secondNumber)
	case "/":
		if secondNumber != 0 {
			result = firstNumber / secondNumber
			fmt.Println("/", firstNumber/secondNumber)
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

// checkAnswer takes the user's input, evaluates it, and compares it to the final result.
func checkAnswer(finalResult int, answer string) bool {
	// Evaluate the user's input as a simple mathematical expression
	evaluatedResult, err := evaluateExpression(answer)
	if err != nil {
		fmt.Println("Error evaluating answer:", err)
		return false
	}
	fmt.Println("Evaluated Result:", evaluatedResult)

	// Check if the evaluated result matches the final result
	return evaluatedResult == finalResult
}

// evaluateExpression takes a string input and evaluates it as a simple arithmetic expression.
func evaluateExpression(expression string) (int, error) {
	var stack []int
	var currentNumber int
	var currentOperation rune = '+'

	for i, char := range expression {
		if unicode.IsDigit(char) {
			// Convert character to integer
			digit, _ := strconv.Atoi(string(char))
			currentNumber = currentNumber*10 + digit
		}

		if !unicode.IsDigit(char) && !unicode.IsSpace(char) || i == len(expression)-1 {
			// Process the current operation when encountering an operator or at the end
			switch currentOperation {
			case '+':
				stack = append(stack, currentNumber)
			case '-':
				stack = append(stack, -currentNumber)
			case '*':
				stack[len(stack)-1] *= currentNumber
			case '/':
				stack[len(stack)-1] /= currentNumber
			}
			currentOperation = char
			currentNumber = 0
		}
	}

	// Sum up the stack
	result := 0
	for _, num := range stack {
		result += num
	}

	return result, nil
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

	// Get user's answer as a string input
	var userAnswer string
	fmt.Println("Enter your calculation to check if it matches the final result: ")
	fmt.Scanln(&userAnswer)

	// Check if the user's answer matches the final result
	if checkAnswer(finalResult, userAnswer) {
		fmt.Println("Correct! Your calculation matches the final result.")
	} else {
		fmt.Println("Incorrect. Your calculation does not match the final result.")
	}
}
