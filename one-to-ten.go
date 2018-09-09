package main

import (
	"os/exec"
	"os"
	"bufio"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
		sayText := "Hello " + scanner.Text() + "!"
		fmt.Printf("%v\n", sayText)
		exec.Command("say", sayText).Run()
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}