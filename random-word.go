package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"math/rand"
	"strings"
	"time"
)

var (
	separator string
	numberOfWords int
)

func init() {
	flag.IntVar(&numberOfWords, "n", 3, "the number of random words to join")
	flag.StringVar(&separator, "s", "-", "a separator to use when joining words")
}

// TODO: break the random word functionality into windows && unix helpers
func main() {
	if len(os.Args) > 1 {
		checkUsage()
	}

	words, err := readAvailableDictionary()
	if err != nil {
		println("Sorry, some unexpected happening reading your dictionary:")
		println(err.Error())
		os.Exit(2)
	}

	flag.Parse()
	rand.Seed(time.Now().Unix())

	pieces := []string{}
	for i := 0; i < numberOfWords; i++ {
		pieces = append(pieces, words[rand.Int() % len(words)])
	}

	println(strings.Join(pieces, separator))
	return
}

// this will fail horribly on windows
func readAvailableDictionary() (words []string, err error) {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	words = strings.Split(string(bytes), "\n")
	return
}

func checkUsage() {
	if os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf(`
usage: random-word -s [separator] -n [number-of-words]
eg: random-word -s="-" -n=5 # holy-moly-guacamole-oily-strombole

The separator between words defaults to '-'
The number of words printed defaults to 3
`)
		os.Exit(1)
	}
}
