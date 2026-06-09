package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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

		if command == "exit" {
			fmt.Println("[Info] Bye!")
			return
		}

		fmt.Printf("%s %s\n\n", command, data)
	}
}
