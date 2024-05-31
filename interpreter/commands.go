package interpreter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

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

func get(args []string) {
	fmt.Println("get value by id")
	if len(args) > 1 {
		log.Fatalf("To many arguments in input command %v", args)
	}
	file := getDatabaseFile(os.O_RDONLY)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
	var fileLines []string
    for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
    }
	for _, line := range fileLines {
		re := regexp.MustCompile(`^key:([-\w]+),value:([\s\w]+);`)
		record := re.FindAllStringSubmatch(line, -1)
		if args[0] == record[0][1] {
			fmt.Println("Get:", record)
		}
	}
}

func getall() {
	fmt.Println("get all pair id:value")
	file := getDatabaseFile(os.O_RDONLY)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan() {
        fmt.Println(fileScanner.Text())
    }
}

func set(args []string) {
	fmt.Println("set value and print its id")
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
