package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
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
	file, err := os.Open("words.txt")
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
	words := strings.Split(all_words, " ")
	file, err = os.Open("characters.txt")
	if err != nil {
		fmt.Println("An error occured, trying to access the list of the available words")
		return
	}
	bts, err = io.ReadAll(file)
	if err != nil {
		fmt.Println("An error occured, trying to read the list of the available words")
		return
	}
	all_chars := string(bts)
	chars := strings.Split(all_chars, "")

	rand.Seed(time.Now().UnixNano())
	var chosenWords []string = []string{}
	for i := 0; i < wordCount; i += 1 {
		new_word := words[rand.Intn(len(words))]
		chosenWords = append(chosenWords, new_word)
	}
	output := chosenWords[0]
	// stick the words together
	for i := 1; i < len(chosenWords); i++ {
		new_con := ""
		if rand.Intn(2) == 0 || !*sc {
			new_con = strconv.Itoa(rand.Intn(1000))
		} else {
			new_con = chars[rand.Intn(len(chars))]
		}
		output += new_con
		output += chosenWords[i]
	}
	fmt.Println(output)
}

var availableSizes map[string]int = map[string]int{
	"xs": 2, "s": 3, "m": 5, "l": 7, "xl": 15,
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
