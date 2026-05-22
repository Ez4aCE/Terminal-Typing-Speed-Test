package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

func main() {
	sentences := []string{
		"The dynamic workspace requires an extra level of justification before publicizing any quirky or bizarre management methods to the staff",
		"A fast-moving, black quartz sphinx was quickly put on display to amaze the skeptical judges who visited the ancient exhibits",
		"Five or six big jet planes zoomed quickly across the bright blue sky, creating a dazzling view for the excited crowd",
		"Back in the quiet valley, several enthusiastic workers built a large frame structure using high-quality zinc parts and heavy copper pipes",
		"A wizardly old man quickly packed all twenty-six heavy boxes with premium fuel, organic juice, and rare glittering quartz crystals",
		"The brave and skillful astronaut quickly maximized his tracking speed to avoid the dangerous space debris flying past the galaxy",
		"Traveling down the rocky river path, our eccentric group realized that they had to quickly bypass several frozen junctions without panic",
		"Crafting a brilliant marketing strategy requires excellent judgment, a massive budget, a few quirky ideas, and zero fear of failure",
		"Cozy sphinx cat breeds often look incredibly unique when they jump quickly into soft velvet chairs for a warm afternoon nap",
		"The lazy delivery driver was heavily criticized after he dropped an extremely valuable crate containing fine Chinese porcelain and quartz clocks",
	}

	sample := sentences[rand.IntN(len(sentences))]
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
