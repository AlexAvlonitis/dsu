package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		show_menu()

		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("1", text) == 0 {
			add_note(*reader)
		}

		if strings.Compare("q", text) == 0 {
			fmt.Println("Goodbye...")
			os.Exit(1)
		}
	}
}

// Display main menu
func show_menu() {
	fmt.Println("")
	fmt.Println(Info("Daily Standup Notes:"))
	fmt.Println("---------------------")
	fmt.Println(Succ("(1)") + " Add a note")
	fmt.Println(Succ("(2)") + " Fetch a note")
	fmt.Println(Succ("(3)") + " Delete a note")
	fmt.Println(Succ("(q)") + " Exit")
	fmt.Println("---------------------")
}
