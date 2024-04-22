package interpreter

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func getDatabaseFile(appendMode int) *os.File {
	filePath, err := os.Getwd()
	if err != nil {
        log.Fatalf("Error with load database file path %v", err)
    }
	file, err := os.OpenFile(filePath + "/interpreter/database-file.txt", appendMode, 0600)
    if err != nil {
        log.Fatalf("Error with load database file %v", err)
    }
	return file
}

func get() {
	fmt.Println("get value by id")
}

func getall() {
	fmt.Println("get all pair id:value")
}

func set(args []string) {
	fmt.Println("set value and return its id")
	fmt.Println("set args:", args)
	file := getDatabaseFile(os.O_APPEND)
	defer file.Close()
	//https://github.com/google/uuid
	for i := range args {
		n, err := file.WriteString(fmt.Sprintf("\nkey:%s,value:%s;", uuid.New().String(), args[i]))
		fmt.Println("is written:", n)
		if err != nil {
			log.Fatalf("Error with writing in file %v", err)
		}
	}
}
