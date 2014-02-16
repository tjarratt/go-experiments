package main

import (
	"flag"
	"fmt"
	"os"
	"math/rand"
	"strings"
)

separator := "-"
numberOfWords := 3
func init() {
	flag.IntVar("n", &numberOfWords)
	flag.StringVar("s", &separator)
}

func main() {
	if os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf(`
usage: random-word -s [separator] -n [number-of-words]
eg: random-word -s="-" -n=5 # holy-moly-guacamole-oily-strombole

The separator between words defaults to '-'
The number of words printed defaults to 3
`)
		os.Exit(1)
	}

	flag.Parse()
	pieces := []string{}
	words = []string{"holy", "moly", "guacamole", "oily", "strombole"}
	for i := 0; i < numberOfWords {
		pieces = append(pieces, words[rand.Int() % len(words)])
	}

	println(strings.Join(pieces, separator))
	return
}
