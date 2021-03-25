package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Prepare to create a note file with a date
func add_note(reader bufio.Reader) {
	for {
		create_home_dir_logs_storage()
		show_add_note_menu()

		fmt.Print("-> ")
		text := getInput(reader)

		if strings.Compare("1", text) == 0 {
			add(reader)
		}

		if strings.Compare("q", text) == 0 {
			break
		}
	}
}

// Display add note menu text
func show_add_note_menu() {
	fmt.Println("")
	fmt.Println(Info("Add a note:"))
	fmt.Println("---------------------")
	fmt.Println(Succ("(1)") + " Add a note for today")
	fmt.Println(Succ("(2)") + " Add a note for another day")
	fmt.Println(Succ("(q)") + " Back")
	fmt.Println("---------------------")
}

// Add a note, creates a filename with the given date as a name
// and it is stored under the logs folder
func add(reader bufio.Reader) {
	fmt.Print(Warn("[Add note +]-> "))
	note := getInput(reader)

	currentTime := time.Now()
	time := currentTime.Format(layoutDate)
	filepath := fmt.Sprintf(logPath() + "/" + time)
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)

	defer f.Close()

	_, err = f.WriteString("• " + note + "\n")
	check(err)

	f.Sync()
	readNote(time)
}

// Create a folder that holds all the logs in the home directory
func create_home_dir_logs_storage() {
	os.Mkdir(logPath(), 0700)
}
