package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Printf("Hello %v!\n", scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}