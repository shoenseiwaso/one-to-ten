package main

import (
	"os/exec"
	"os"
	"bufio"
	"fmt"
)

var scanner *bufio.Scanner

func outLine(line string) {
	fmt.Printf("%v\n", line)
	exec.Command("say", line).Run()	
}

func inLine() string {
	succeeded := scanner.Scan()
	if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	
	if succeeded {
		return scanner.Text()
	}

	return ""
}

func main() {
	// setup
	scanner = bufio.NewScanner(os.Stdin)
    
	outLine("2 + 2")
	line := inLine()
	outLine(line)
}