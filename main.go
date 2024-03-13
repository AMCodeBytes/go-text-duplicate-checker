package main

import (
	"bufio"
	"fmt"
	"os"
)

var textLines map[string]bool = make(map[string]bool)

func main() {
	var choice string
	fmt.Println("---  Text File Duplicate Checker  ---")

	for {
		fmt.Println("check: check for duplicates")
		fmt.Println("quit: leave the application")
		fmt.Print("Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case "check":
			check()
		case "quit":
			os.Exit(0)
		}
	}
}

func check() {
	var fileName string = ""
	var fileLines []string

	readFile, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	defer readFile.Close()

	fmt.Print("Enter name of text file to append: ")
	fmt.Scan(&fileName)

	file, err := os.OpenFile(fileName+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Filter the lines from the read file
	for _, line := range fileLines {
		// fmt.Println(line)
		if !duplicateCheck(line) {
			addLine(line)
		}
	}

	fmt.Println(textLines)

	// Insert the filtered lines into a new file
	for value, ok := range textLines {
		if ok {
			_, err2 := file.WriteString(value + "\n")
			if err2 != nil {
				fmt.Println("Could not write text to " + fileName + ".txt")
			} else {
				fmt.Println("Operation successful! Text has been appended to " + fileName + ".txt")
			}
		}
	}
}

func duplicateCheck(text string) bool {
	_, yes := textLines[text]

	return yes
}

func addLine(line string) {
	textLines[line] = true
}
