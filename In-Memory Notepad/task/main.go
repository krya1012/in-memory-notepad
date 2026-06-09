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

		case "update":
			if data == "" {
				fmt.Println("[Error] Missing position argument")
				break
			}
			parts := strings.SplitN(data, " ", 2)
			posStr := parts[0]
			noteStr := ""
			if len(parts) > 1 {
				noteStr = parts[1]
			}
			pos, err := strconv.Atoi(posStr)
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", posStr)
				break
			}
			if pos < 1 || pos > maxNotes {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", pos, maxNotes)
				break
			}
			if pos > len(notes) {
				fmt.Println("[Error] There is nothing to update")
				break
			}
			if strings.TrimSpace(noteStr) == "" {
				fmt.Println("[Error] Missing note argument")
				break
			}
			notes[pos-1] = noteStr
			fmt.Printf("[OK] The note at position %d was successfully updated\n", pos)

		case "delete":
			if strings.TrimSpace(data) == "" {
				fmt.Println("[Error] Missing position argument")
				break
			}
			pos, err := strconv.Atoi(strings.TrimSpace(data))
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", strings.TrimSpace(data))
				break
			}
			if pos < 1 || pos > maxNotes {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", pos, maxNotes)
				break
			}
			if pos > len(notes) {
				fmt.Println("[Error] There is nothing to delete")
				break
			}
			notes = append(notes[:pos-1], notes[pos:]...)
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", pos)

		case "exit":
			fmt.Println("[Info] Bye!")
			return

		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}
