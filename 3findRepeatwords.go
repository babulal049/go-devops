package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readsentence := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	sentence, err := readsentence.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	fmt.Print("You entered:", sentence)
	findRepeatingWords(sentence)

}

func findRepeatingWords(inputString string) string {
	words := strings.Fields(inputString)
	wordCount := make(map[string]int)
	for _, word := range words {
		cleanedWord := strings.ToLower(strings.Trim(word, ".,")) // Normalize words
		wordCount[cleanedWord]++
	}

	// Check for repeated words
	fmt.Println("Repeated words:")
	for word, count := range wordCount {
		if count > 1 {
			fmt.Println(word)
		}
	}

	return inputString
}
