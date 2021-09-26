package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type word struct {
	name        string
	description string
}

func newWord(n string, d string) word {
	w := word{
		name:        n,
		description: "",
	}

	return w
}

func readWordFromFile() []word {

	file, err := os.Open("wordList.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []word

	for scanner.Scan() {
		w := scanner.Text()
		wordSplit := strings.Split(w, ":")

		words = append(words, word{name: wordSplit[0], description: wordSplit[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

func (wr *word) shuffleLetters(w word) string {

	rand.Seed(time.Now().Unix())

	shuffledWord := []rune(w.name)
	rand.Shuffle(len(shuffledWord), func(i, j int) {
		shuffledWord[i], shuffledWord[j] = shuffledWord[j], shuffledWord[i]
	})

	sw := string(shuffledWord)

	return sw
}
