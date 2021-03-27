package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Prepare to delete a note file with a date
func delete_note(reader bufio.Reader) {
	for {
		show_delete_note_menu()

		fmt.Print("-> ")
		text := getInput(reader)

		if strings.Compare("1", text) == 0 {
			delete(reader)
		}

		if strings.Compare("q", text) == 0 {
			break
		}
	}
}

// Display delete note menu text
func show_delete_note_menu() {
	fmt.Println("")
	fmt.Println(Info("Delete a note:"))
	fmt.Println("---------------------")
	fmt.Println(Succ("(1)") + " Delete all the notes for a day")
	fmt.Println(Succ("(2)") + " Delete a line from a note")
	fmt.Println(Succ("(q)") + " Back")
	fmt.Println("---------------------")
}

// delete the file that stores the notes for the given date
func delete(reader bufio.Reader) {
	fmt.Print("Enter date to delete (yy-mm-dd) -> ")
	text := getInput(reader)

	e := os.Remove(logPath() + "/" + text)
	if e != nil {
		log.Fatal(e)
	} else {
		fmt.Println(Succ("\nNotes for " + text + " deleted"))
	}
}