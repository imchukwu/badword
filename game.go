package main

import (
	"math/rand"
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

// func readWordFromFile() []word {

// }

func (wr *word) shuffleLetters(w word) string {

	rand.Seed(time.Now().Unix())

	shuffledWord := []rune(w.name)
	rand.Shuffle(len(shuffledWord), func(i, j int) {
		shuffledWord[i], shuffledWord[j] = shuffledWord[j], shuffledWord[i]
	})

	sw := string(shuffledWord)

	return sw
}
