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

func getAnswer(w word) {
	total := 0
	reader := bufio.NewReader(os.Stdin)
	username, _ := getInput("Enter your username: ", reader)

	badWord := w.shuffleLetters(w)

	fmt.Println("Shuffled Word is: ", badWord)

	correctWord, _ := getInput("Provide the correct word: ", reader)

	result := correctWord == w.name

	if result != true {
		fmt.Printf("%v got the answer wrongly with total score %v", username, total)
	} else {
		total += 1
		fmt.Printf("%v got the answer correct with total score %v", username, total)
	}
}

func main() {
	getAnswer(newWord("tiger", "wild animal"))

	// w := newWord()
	// word := w.shuffleLetters(w)

	// fmt.Println(word)

	// file, err := os.Open("wordList.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// var words []word

	// for scanner.Scan() {
	// 	w := strings.Split(scanner.Text(), "\t")

	// 	// fmt.Println(w)
	// 	words = append(words, word{value: w[0], description: w[0]})
	// 	fmt.Println(w[0])
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(words)

}
