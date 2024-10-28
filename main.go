package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CalculationResult struct {
	Result int
	Steps  string
}

type GameData struct {
	Target        int
	ChosenNumbers []int
	Results       []CalculationResult
	UserResult    int
	Difference    int
}

func chooseBigNumbers(numberOfBigNumber int) []int {
	bigNumbers := []int{25, 50, 75, 100}
	numbersChosen := []int{}

	for i := 0; i < numberOfBigNumber; i++ {
		randomIndex := rand.Intn(len(bigNumbers))
		randomElement := bigNumbers[randomIndex]
		numbersChosen = append(numbersChosen, randomElement)
		bigNumbers = append(bigNumbers[:randomIndex], bigNumbers[randomIndex+1:]...)
	}

	return numbersChosen
}

func chooseSmallNumbers(allNumbersChosen []int) []int {
	for len(allNumbersChosen) < 6 {
		randomNumber := rand.Intn(10) + 1
		allNumbersChosen = append(allNumbersChosen, randomNumber)
	}
	return allNumbersChosen
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

func applyOperations(numbers []int, currentResult int, index int, operations []string, target int, steps string, validResults *[]CalculationResult) {
	if index == len(numbers) {
		if abs(currentResult-target) <= 10 {
			*validResults = append(*validResults, CalculationResult{Result: currentResult, Steps: steps})
		}
		return
	}

	for _, op := range operations {
		nextResult := doCalculation(currentResult, numbers[index], op)
		newSteps := fmt.Sprintf("%s %s %d", steps, op, numbers[index])
		applyOperations(numbers, nextResult, index+1, operations, target, newSteps, validResults)
	}
}

func calculateOnCombinations(allNumbersChosen []int, target int) []CalculationResult {
	operations := []string{"+", "-", "*", "/"}
	var validResults []CalculationResult

	for size := 2; size <= len(allNumbersChosen); size++ {
		combinations := generateCombinations(allNumbersChosen, size)
		for _, combo := range combinations {
			initialStep := fmt.Sprintf("%d", combo[0])
			applyOperations(combo, combo[0], 1, operations, target, initialStep, &validResults)
		}
	}
	return validResults
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	numberOfBigNumber, _ := strconv.Atoi(r.URL.Query().Get("bigNumber"))

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(900) + 100
	allNumbersChosen := chooseBigNumbers(numberOfBigNumber)
	allNumbersChosen = chooseSmallNumbers(allNumbersChosen)
	validResults := calculateOnCombinations(allNumbersChosen, target)

	data := GameData{
		Target:        target,
		ChosenNumbers: allNumbersChosen,
		Results:       validResults,
	}

	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, data)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	userResult, _ := strconv.Atoi(r.FormValue("userSolution"))
	target, _ := strconv.Atoi(r.FormValue("target"))
	difference := abs(userResult - target)

	// Recreate the game data to show results
	chosenNumbersStr := r.FormValue("chosenNumbers")
	chosenNumbers := parseChosenNumbers(chosenNumbersStr)
	validResults := calculateOnCombinations(chosenNumbers, target)

	data := GameData{
		Target:        target,
		ChosenNumbers: chosenNumbers,
		Results:       validResults,
		UserResult:    userResult,
		Difference:    difference,
	}

	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, data)
}

func parseChosenNumbers(s string) []int {
	parts := strings.Split(strings.Trim(s, "[]"), " ")
	var numbers []int
	for _, part := range parts {
		if n, err := strconv.Atoi(part); err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func main() {
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/submit", submitHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
