package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

// TODO: how to handle lots of printlns that accumulate over time
func main() {
	println("starting REPL...")

	file := createTempfile()

	userInputHistory := []string{}

	for {
		file.Seek(0, 0)

		userInput := readInput()
		print(" " + userInput)

		userInputHistory = append(userInputHistory, userInput)

		writeTempfileHeader(file)
		updateTempfile(userInputHistory, file)
		writeTempfileFooter(file)

		_, formatErr := exec.Command("go", "fmt", file.Name()).Output()
		if formatErr != nil {
			println(fmt.Sprintf("unexpected error formatting: %s", formatErr.Error()))
			os.Exit(1)
		}

		cmd := exec.Command("go", "run", file.Name())
		out, err := cmd.Output()

		if err != nil {
			println(fmt.Sprintf("Compiler is PISSED. Error: %s", err.Error()))
			userInputHistory = userInputHistory[:len(userInputHistory)-1]

			// TODO: consider falling back to the last known good compilation state41
			// TODO: think about multi-line (open curly, etc)
			// TODO: think about adding / removing import
		} else {
			fmt.Print(string(bytes.Join(regexp.MustCompile(".*s\n").FindAll(out, -1), nil)))
			println(fmt.Sprintf("we read %d characters of output: '%s'", len(out), out))
		}
	}
}

func readInput() string {
	print("gREPL: > ")
	bio := bufio.NewReader(os.Stdin)
	userInput, _ := bio.ReadString('\n')

	return userInput
}

func createTempfile() *os.File {
	file, err := ioutil.TempFile("/tmp", "go-repl")
	if err != nil {
		println("whelp, couldn't create a tempfile. BAILING!")
		os.Exit(1)
	}

	file.Close()

	properName := file.Name() + ".go"
	mvCmd := exec.Command("mv", file.Name(), properName)
	_, mvErr := mvCmd.Output()

	if mvErr != nil {
		println(fmt.Sprintf("could not move tempfile %s to %s. Bailing.", file.Name(), properName))
		os.Exit(1)
	}

	renamedFile, err := os.OpenFile(properName, os.O_RDWR, 0)
	if err != nil {
		println(fmt.Sprintf("could not open tempfile %s. Bailing.", properName))
		os.Exit(1)
	}

	return renamedFile
}

func updateTempfile(lines []string, filehandle *os.File) {
	for _, line := range lines {
		status, err := filehandle.WriteString(line)

		if err != nil {
			println("shit, couldn't append to file. BAILING!")
			println(fmt.Sprintf("%#v, %s", status, err.Error()))
			os.Exit(1)
		}
	}
}

// TODO: keep track of packages we might need to import
// should be fast enough on error to just re-compile
// with missing packages / without extra imports
func writeTempfileHeader(file *os.File) {
	file.WriteString("package main\n func main() {\n")
}

func writeTempfileFooter(file *os.File) {
	file.WriteString("\n}\n")
}
