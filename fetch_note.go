package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

// Prepare to fetch a note by date
func fetch_note(reader bufio.Reader) {
	for {
		show_fetch_note_menu()

		fmt.Print("-> ")
		text := getInput(reader)

		if strings.Compare("1", text) == 0 {
			fetchLastNote()
		}

		if strings.Compare("q", text) == 0 {
			break
		}
	}
}

// Display fetch note menu
func show_fetch_note_menu() {
	fmt.Println("")
	fmt.Println(Info("Fetch a note:"))
	fmt.Println("---------------------")
	fmt.Println(Succ("(1)") + " Fetch last notes")
	fmt.Println(Succ("(2)") + " Fetch second from last notes")
	fmt.Println(Succ("(3)") + " Fetch notes from a specific date")
	fmt.Println(Succ("(4)") + " Show all available notes")
	fmt.Println(Succ("(q)") + " Back")
	fmt.Println("---------------------")
}

// Parse all notes in the logs directory and display the most recent one
func fetchLastNote() {
	fileInfo, err := ioutil.ReadDir(logPath())
	check(err)

	var files []string
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	if len(files) > 0 {
		lastNote := files[0]
		readNote(lastNote)
		return
	}

	fmt.Println("There are no available notes")
	enterToContinue()
}

// read a specific note file from the logs folder
func readNote(noteName string) {
	content, err := ioutil.ReadFile(logPath() + "/" + noteName)
	check(err)

	fmt.Printf(Teal("\nFile contents [%s]:\n \n%s"), noteName, content)
	enterToContinue()
}
