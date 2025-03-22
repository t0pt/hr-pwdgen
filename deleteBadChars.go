package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("words_alpha.txt")
	if err != nil {
		fmt.Println("An error occured, trying to access the list of the available words")
		return
	}
	bts, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("An error occured, trying to read the list of the available words")
		return
	}
	all_words := string(bts)

	good_words := removeUnseenCharacters(all_words)
	err = os.WriteFile("words.txt", []byte(good_words), 0644)
	if err != nil {
		panic(err)
	}
}

func removeUnseenCharacters(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsPrint(r) { // Keep only printable characters
			builder.WriteRune(r)
		} else if r == '\n' {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}
