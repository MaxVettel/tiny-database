package interpreter

import (
	"fmt"
	"os"
)

func HandleCommand(text string) {
	switch text {
	case "stop":
		os.Exit(0)
	case "get":
		get()
	case "getall":
		getall()
	case "set":
		set()
	default:
		fmt.Println("unknown command")
	}
}
