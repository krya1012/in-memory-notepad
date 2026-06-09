package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the maximum number of notes: ")
	scanner.Scan()
	maxNotes, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	var notes []string

	for {
		fmt.Print("\nEnter a command and data: ")
		scanner.Scan()
		input := strings.SplitN(scanner.Text(), " ", 2)
		command := input[0]
		data := ""
		if len(input) > 1 {
			data = input[1]
		}

		switch command {
		case "create":
			if strings.TrimSpace(data) == "" {
				fmt.Println("[Error] Missing note argument")
			} else if len(notes) >= maxNotes {
				fmt.Println("[Error] Notepad is full")
			} else {
				notes = append(notes, data)
				fmt.Println("[OK] The note was successfully created")
			}
		case "list":
			if len(notes) == 0 {
				fmt.Println("[Info] Notepad is empty")
			} else {
				for i, note := range notes {
					fmt.Printf("[Info] %d: %s\n", i+1, note)
				}
			}
		case "clear":
			notes = notes[:0]
			fmt.Println("[OK] All notes were successfully deleted")
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}
