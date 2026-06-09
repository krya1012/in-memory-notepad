package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxNotes = 5

func main() {
	var notes []string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a command and data: ")
		scanner.Scan()
		input := strings.SplitN(scanner.Text(), " ", 2)
		command := input[0]
		data := ""
		if len(input) > 1 {
			data = input[1]
		}

		switch command {
		case "create":
			if len(notes) >= maxNotes {
				fmt.Println("[Error] Notepad is full")
			} else {
				notes = append(notes, data)
				fmt.Println("[OK] The note was successfully created")
			}
		case "list":
			for i, note := range notes {
				fmt.Printf("[Info] %d: %s\n", i+1, note)
			}
		case "clear":
			notes = notes[:0]
			fmt.Println("[OK] All notes were successfully deleted")
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Printf("%s %s\n", command, data)
		}
		fmt.Println()
	}
}
