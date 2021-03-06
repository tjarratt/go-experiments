package main

import (
	"bufio"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		println("usage: used_imports path/to/some/file.go")
		os.Exit(1)
	}

	cmd := exec.Command("go", "build", os.Args[1])
	stderr, err := cmd.StderrPipe()
	if err != nil {
		println("fatal error getting stderr pipe")
		os.Exit(2)
	}

	cmd.Start()
	r := bufio.NewReader(stderr)
	importRegex := regexp.MustCompile("imported and not used: \"(.*)\"")
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		matches := importRegex.FindAllStringSubmatch(string(line), 1)
		if matches == nil {
			continue
		}

		println(matches[0][1])
	}
	return
}
