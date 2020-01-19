package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

const PYTHON_FILE_NAME = "neopixel-interface.py"

func main() {
	cmd := exec.Command("python3", PYTHON_FILE_NAME)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)
	cmd.Wait()
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}