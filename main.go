package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func getAnswer(w []word) {
	count := 0
	total := 0
	reader := bufio.NewReader(os.Stdin)
	username, _ := getInput("Enter your username: ", reader)
	fmt.Print("\n")

	for _, v := range w {
		count++
		badWord := v.shuffleLetters(v)

		fmt.Println("Word:", count)
		fmt.Println("Shuffled Word is: ", strings.ToUpper(badWord))
		fmt.Println("Hint: ", v.description)

		correctWord, _ := getInput("Provide the correct word: ", reader)

		result := strings.ToUpper(correctWord) == strings.ToUpper(v.name)

		if result != true {
			fmt.Println("The answer is wrong")
		} else {
			total++
			fmt.Println("The answer is Correct")
		}
		fmt.Print("\n")

	}

	fmt.Println(username, "have a total score of:", total, "of:", count, "words")
}

func main() {
	getAnswer(readWordFromFile())

}
