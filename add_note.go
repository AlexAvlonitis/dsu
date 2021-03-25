package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
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
			fmt.Print(Warn("[Add note +]-> "))
			note := getInput(reader)

			currentTime := time.Now()
			formattedTime := currentTime.Format(layoutDate)
			add(note, formattedTime)
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
func add(note string, time string) {
	usr, err := user.Current()
	check(err)

	filepath := fmt.Sprintf("%s/dsu_logs/%s", usr.HomeDir, time)
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
