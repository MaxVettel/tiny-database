package interpreter

import (
	"fmt"
	"os"
	"regexp"
)

func HandleCommand(text string) {
	re := regexp.MustCompile(`("[^"]+"|[^\s"]+)`)
	args := re.FindAllString(text, -1)
	fmt.Println("Command:", args[0])
	fmt.Println("Args:", args[1:])
	switch args[0] {
	case "stop":
		os.Exit(0)
	case "get":
		get(args[1:])
	case "getall":
		getall()
	case "set":
		set(args[1:])
	case "delete":
		delete(args[1:])
	case "":
		fmt.Println("empty command")
	default:
		fmt.Println("unknown command")
	}
}
