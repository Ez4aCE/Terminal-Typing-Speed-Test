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

	duration := end.Sub(start)

	words := strings.Fields(sample)
	wordCount := len(words)

	minutes := duration.Minutes()

	wpm := float64(wordCount) / minutes
	fmt.Printf("Time taken : %v \n", duration.Seconds())
	fmt.Println("You typed : ", typed)

	fmt.Printf("Gross WPM : %.2f \n", wpm)

}
