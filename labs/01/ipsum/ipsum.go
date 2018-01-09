package main

import (
	"fmt"
	"math/rand"
	"time"
)

var wordbank = []string{
	"foo",
	"bar",
	"baz",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generateIpsum(100, 6))
}

func generateIpsum(wordCount, sentenceLength int) string {
	ipsum := ""
	for wordsLeft := wordCount; wordsLeft > 0; wordsLeft -= sentenceLength {
		numWords := sentenceLength
		if wordsLeft < numWords {
			numWords = wordsLeft
		}
		if ipsum != "" {
			ipsum += "  "
		}
		ipsum += generateSentence(numWords)
	}
	return ipsum
}

func generateSentence(wordCount int) string {
<<<<<<< HEAD
	ipsum := ""
	for i := 0; i < wordCount; i++ {
		if ipsum != "" {
			ipsum += " "
		}
		ipsum += getWord()
	}
	return ipsum + "."
=======

>>>>>>> 41b0aaa71f904b1262386f81eb52e4d0a14378b3
}

func getWord() string {
	index := rand.Intn(len(wordbank))
	return wordbank[index]
}
