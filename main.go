package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Printf("Time taken : %v \n", duration)
	fmt.Println("You typed : ", typed)

}
