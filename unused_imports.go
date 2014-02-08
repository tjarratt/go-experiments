package main

import(
	"bufio"
	"os"
  "os/exec"
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

	err = cmd.Run()
	if err == nil {
		println("no unused imports.")
		os.Exit(0)
	}

	line, err := bufio.NewReader(stderr).ReadString('\n')
	if err != nil {
		println("uh oh")
	}
	println("output", line)

	return
}
