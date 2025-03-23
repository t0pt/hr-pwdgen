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
	"unicode"
)

func main() {
	help := flag.Bool("help", false, "help")
	size := flag.String("s", "m", "size of the password (m by default)")
	sc := flag.Bool("sc", false, "disable special special characters (enabled by default)")
	f := flag.Bool("f", false, "fancyfy the words (change some a's to @, some o's to 0 and some i's to 1) (enabled by default)")
	c := flag.Bool("c", false, "capialize some letters in words (disabled by default)")
	hotness := flag.Int("hotness", 5, "probability that the character will be fancifyed: 10:n , where n is the passed number (defaults to 5, so it's 2:1 probability)")
	flag.Parse()
	*sc = !*sc
	*f = !*f

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
		new_word = CapitalizeFirst(new_word)
		if rand.Intn(10+*hotness) < *hotness && *c {
			new_word = CapitalizeSome(new_word, *hotness)
		}
		if rand.Intn(10+*hotness) < *hotness && *f {
			new_word = SpecialCharsString(new_word, *hotness)
		}
		chosenWords = append(chosenWords, new_word)
	}
	output := chosenWords[0]
	// stick the words together
	for i := 1; i < len(chosenWords); i++ {
		new_con := ""
		if rand.Intn(2) == 0 || !*sc {
			new_con = strconv.Itoa(rand.Intn(100))
		} else {
			new_con = chars[rand.Intn(len(chars))]
		}
		output += new_con
		output += chosenWords[i]
	}
	fmt.Println(output)
}

func SpecialCharsString(s string, h int) string {
	output := ""
	for _, run := range s {
		switch run {
		case 'a':
			if rand.Intn(10+h) < h {
				run = '@'
			}
		case 'o':
			if rand.Intn(10+h) < h {
				run = '0'
			}
		case 'i':
			if rand.Intn(10+h) < h {
				run = '1'
			}
		}
		output += string(run)
	}
	return output
}

func CapitalizeSome(s string, h int) string {
	output := ""
	for _, run := range s {
		if rand.Intn(10+h) < h {
			run = unicode.ToUpper(run)
		}
		output += string(run)
	}
	return output
}

func CapitalizeFirst(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
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
  -help         help
  -s            size of the password (m by default) (default "m")
  -sc           disable special special characters (enabled by default)
  -f            fancyfy the words (change some a's to @, some o's to 0 and some i's to 1) (enabled by default)
  -c            capialize some letters in words (disabled by default)
  -hotness      probability that the character will be fancifyed: 10:n , where n is the passed number (defaults to 5, so it's 2:1 probability)

Password size table:
%s`, sizeTable)
