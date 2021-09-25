package main

import (
	"fmt"
)

func main() {

	w := newWord()
	word := w.shuffleLetters(w)

	fmt.Println(word)

}
