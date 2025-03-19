package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	help := flag.Bool("h", false, "help")
	size := flag.String("s", "m", "size of the password (m by default)")
	sc := flag.Bool("sc", false, "disable special special characters (enabled by default)")
	flag.Parse()
	*sc = !*sc

	if *help {
		fmt.Println(helpMessage)
		return
	}

	wordCount := int(0)
	for key, val := range availableSizes {
		if key == *size {
			wordCount = val
			break
		}
	}
	if wordCount == 0 {
		fmt.Println("bad size provided, please choose one from the following:")
		fmt.Println(sizeTable)
		return
	}

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
	words := strings.Split(all_words, "\n")

	rand.Seed(time.Now().UnixNano())
	var foundWords []string = []string{}
	for i := 0; i < wordCount; i += 1 {
		new_word := words[rand.Intn(len(words))]
		fmt.Println(new_word)
		foundWords = append(foundWords, new_word)
	}
	fmt.Println("Words list:", foundWords)
	fmt.Println(strings.Join(foundWords, "."))
}

var availableSizes map[string]int = map[string]int{
	"xs": 3, "s": 5, "m": 7, "l": 10, "xl": 15,
}

var sizeTable string = fmt.Sprintf(`    xs - %d words + random spacers
    s - %d words + random spacers
    m - %d words + random spacers
    l - %d words + random spacers
    xl - %d words + random spacers`, availableSizes["xs"], availableSizes["s"],
	availableSizes["m"], availableSizes["l"], availableSizes["xl"])

var helpMessage string = fmt.Sprintf(`Usage:
  -h    help
  -s string
        size of the password (default "m")
  -sc
        disable special special charecters (enabled by default)

Size table:
%s`, sizeTable)
