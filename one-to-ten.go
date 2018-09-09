package main

import (
	"strconv"
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
	
	// introduction
	outLine("Welcome to the one to ten game. Good luck!")
	outLine("What is your name?")
	name := inLine()
	outLine("Let's begin " + name)
	points := 0

	// game loop
	for a := 1; a <= 21; a++ {
		outLine(fmt.Sprintf("What is %v + %v?", a, a))
		answer := inLine()
		rightAnswer := a + a

		number, err := strconv.Atoi(answer)
		if err != nil {
			outLine("That's not a number, " + name + ", please enter a number!")
			continue
		}

		if number == rightAnswer {
			outLine("Good job " + name + ", that is correct!")
			points++
		} else {
			outLine(fmt.Sprintf("Sorry, %v, that is the wrong answer. The right answer is %v", name, rightAnswer))
		}
	}

	// print result
	outLine(fmt.Sprintf("Congratulations %v you scored %v points!", name, points))
}