package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	sample := "The quick brown fox jumps over the lazy dog"
	fmt.Println("Type the following sentence as fast as you can : ")
	fmt.Println(sample)

	start := time.Now()

	reader := bufio.NewReader(os.Stdin)
	typed, _ := reader.ReadString('\n')
	end := time.Now()
	typed = strings.TrimSuffix(typed, "\n")
	duration := end.Sub(start)

	sampleWords := strings.Fields(sample)
	wordCount := len(sampleWords)

	minutes := duration.Minutes()

	grossWPM := float64(wordCount) / minutes

	typedWords := strings.Fields(typed)
	correct := 0
	for i := 0; i < len(typedWords) && i < len(sampleWords); i++ {
		if typedWords[i] == sampleWords[i] {
			correct++
		}

	}
	accuracy := (float64(correct) / float64(wordCount)) * 100

	netWPM := (accuracy * grossWPM) / 100

	fmt.Printf("Time taken : %v \n", duration.Seconds())
	fmt.Println("You typed : ", typed)
	fmt.Printf("Gross WPM : %.2f \n", grossWPM)
	fmt.Printf("Accuracy : %.2f \n", accuracy)
	fmt.Printf("Net WPM : %.2f \n", netWPM)
}
