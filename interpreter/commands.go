package interpreter

import (
	"fmt"
	"log"
	"os"
)

func get() {
	fmt.Println("get value by id")
}

func getall() {
	fmt.Println("get all pair id:value")
	filePath, err := os.Getwd()
	if err != nil {
        log.Fatalf("Error with load database file path %v", err)
    }
	file, err := os.ReadFile(filePath + "/interpreter/database-file.txt")
    if err != nil {
        log.Fatalf("Error with load database file %v", err)
    }
	fmt.Println(string(file))
}

func set() {
	fmt.Println("set value and return its id")
}
