//Function to generate all permutations to get to that number
//Print out all permutations

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
		result = firstNumber
	}
	return result
}

// Helper function to generate all permutations of a slice of integers
func permute(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		perm := make([]int, len(nums))
		copy(perm, nums)
		*result = append(*result, perm)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permute(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start] // backtrack
	}
}

// Helper function to apply operations on two integers
func applyOperation(a, b int, op string) (int, bool) {
	switch op {
	case "+":
		return a + b, true
	case "-":
		return a - b, true
	case "*":
		return a * b, true
	case "/":
		if b != 0 && a%b == 0 { // Ensure no division by zero and integer division
			return a / b, true
		}
		return 0, false
	}
	return 0, false
}

// Function to check all possible results for a given permutation of numbers and operations
func evaluatePermutation(nums []int, ops []string, finalResult int) {
	for _, op1 := range ops {
		for _, op2 := range ops {
			for _, op3 := range ops {
				for _, op4 := range ops {
					// Apply the operations between the numbers
					result1, valid := applyOperation(nums[0], nums[1], op1)
					if !valid {
						continue
					}
					result2, valid := applyOperation(result1, nums[2], op2)
					if !valid {
						continue
					}
					result3, valid := applyOperation(result2, nums[3], op3)
					if !valid {
						continue
					}
					result4, valid := applyOperation(result3, nums[4], op4)
					if !valid {
						continue
					}

					// Check if the result matches the finalResult
					if result4 == finalResult {
						fmt.Printf("%d %s %d %s %d %s %d %s %d = %d\n", nums[0], op1, nums[1], op2, nums[2], op3, nums[3], op4, nums[4], finalResult)
					}
				}
			}
		}
	}
}

// Main function to generate permutations and evaluate them
func findPermutationsAndEvaluate(finalResult int, numbers []int) {
	ops := []string{"+", "-", "*", "/"}

	// Generate all permutations of the numbers
	var permutations [][]int
	permute(numbers, 0, &permutations)

	// Evaluate each permutation with the operations
	for _, perm := range permutations {
		evaluatePermutation(perm, ops, finalResult)
	}
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

// evaluateExpression takes a string input and evaluates it as a simple arithmetic expression. NAUGHTY
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

	findPermutationsAndEvaluate(finalResult, allNumbers)

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
