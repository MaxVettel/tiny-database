package main

import (
	"bufio"
	"fmt"
	"os"
	"tiny-database/interpreter"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		interpreter.HandleCommand(text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
