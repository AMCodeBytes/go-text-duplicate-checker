package main

import (
	"fmt"
	"os"
)

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
	fmt.Println("Check")
}
