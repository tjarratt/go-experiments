package main

import (
	"os"
	"bufio"
	"os/exec"
	"io/ioutil"
)

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

		cmd := exec.Command("/usr/local/go/bin/go", "run", file.Name())
		out, err := cmd.Output()
		if err != nil {
			println("compiler is PISSED")
			println(err.Error())

			// fall back, removing this line?
			// TODO: think about multi-line
			// TODO: think about adding / removing import statements?
		}

		println(string(out))
	}
	// TODO: how to handle lots of printlns that accumulate over time
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

	return file
}

func updateTempfile(lines []string, filehandle *os.File) {
	for _, line := range lines {
		status, err := filehandle.WriteString(line)

		if err != nil {
			println("shit, couldn't append to file. BAILING!")
			println(status, err)
			os.Exit(1)
		}
	}
}

func writeTempfileHeader(file *os.File) {
	file.WriteString("package main\n func main() {\n")
}

func writeTempfileFooter(file *os.File) {
	file.WriteString("\n}\n")
}
