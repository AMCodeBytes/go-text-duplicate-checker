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
	// data, err := os.ReadFile("test.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	readFile, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	defer readFile.Close()

	for _, line := range fileLines {
		// fmt.Println(line)
		if !duplicateCheck(line) {
			addLine(line)
		}
	}

	fmt.Println(textLines)
}

func duplicateCheck(text string) bool {
	_, yes := textLines[text]

	return yes
}

func addLine(line string) {
	textLines[line] = true
}
