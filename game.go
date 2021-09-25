package main

import (
	"math/rand"
	"time"
)

type word struct {
	value       string
	description string
}

func newWord() word {
	w := word{
		value:       "cat",
		description: "",
	}

	return w
}

func (wr *word) shuffleLetters(w word) string {

	rand.Seed(time.Now().Unix())

	shuffledWord := []rune(w.value)
	rand.Shuffle(len(shuffledWord), func(i, j int) {
		shuffledWord[i], shuffledWord[j] = shuffledWord[j], shuffledWord[i]
	})

	sw := string(shuffledWord)

	return sw
}
