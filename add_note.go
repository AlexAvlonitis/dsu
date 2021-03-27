package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			currentTime := time.Now()
			time := currentTime.Format(layoutDate)
			add(time, reader)
		}

		if strings.Compare("2", text) == 0 {
			fmt.Print("Enter date (yy-mm-dd) ->")
			text := getInput(reader)
			add(text, reader)
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
	fmt.Println(Succ("(1)") + " Add today's notes")
	fmt.Println(Succ("(2)") + " Add a note for another day")
	fmt.Println(Succ("(q)") + " Back")
	fmt.Println("---------------------")
}

// Add a note, creates a filename or appends if the file exists,
// with a date param as a filename and stores it under the logs folder.
func add(time string, reader bufio.Reader) {
	fmt.Print(Warn("[Add note +]-> "))
	note := getInput(reader)

	// open/create the file
	filepath := fmt.Sprintf(logPath() + "/" + time)
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	// get next number of line and write to file
	nextNumber := LinesInFile(f) + 1
	n := strconv.Itoa(nextNumber)
	_, err = f.WriteString(n + ") " + note + "\n")
	check(err)

	f.Sync()
	readNote(time)
}

// reads the number of lines of a file
func LinesInFile(f *os.File) int {
	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return len(result)
}

// Create a folder that holds all the logs in the home directory
func create_home_dir_logs_storage() {
	os.Mkdir(logPath(), 0700)
}
