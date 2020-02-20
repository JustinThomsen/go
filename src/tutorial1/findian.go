package main

import "fmt"
import "strings"

func main() {
	var userInput string

	fmt.Println("Enter a string: ")
	count, error := fmt.Scan(&userInput)

	if error != nil {
		fmt.Print(error)
	}
	if count != 0 {
		findIan(userInput)
	}
}

func findIan(userInput string) {
	userInputLower := strings.ToLower(userInput)
	lastIndexOfString := len(userInputLower) - 1
	lastIndexOfN := strings.LastIndex(userInputLower, "n")
	firstIndexOfI := strings.Index(userInputLower, "i")
	containsA := strings.Contains(userInputLower, "a")
	if firstIndexOfI == 0 &&
		containsA &&
		lastIndexOfN == lastIndexOfString {
		fmt.Print("You did it!")
	}
}

