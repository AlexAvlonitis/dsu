package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Prepare to fetch a note by date
func fetch_note(reader bufio.Reader) {
	for {
		show_fetch_note_menu()

		fmt.Print("-> ")
		text := getInput(reader)

		if strings.Compare("1", text) == 0 {
			fetchLastNotes()
		}
		if strings.Compare("2", text) == 0 {
			fetchNotesbyDate(reader)
		}
		if strings.Compare("3", text) == 0 {
			fetchAllNotes()
		}

		if strings.Compare("q", text) == 0 {
			break
		}
		clearScreen()
	}
}

// Display fetch note menu
func show_fetch_note_menu() {
	fmt.Println("")
	fmt.Println(Info("Fetch a note:"))
	fmt.Println("--------------------------------------")
	fmt.Println(Succ("(1)") + " Fetch last 2 notes")
	fmt.Println(Succ("(2)") + " Fetch notes from a specific date")
	fmt.Println(Succ("(3)") + " Show all available notes")
	fmt.Println(Succ("(q)") + " Back")
	fmt.Println("--------------------------------------")
}

// Parse all notes in the logs directory and display the most recent one
// if it's available
func fetchLastNotes() {
	_, files := parseLogsDir()

	if len(files) > 0 {
		lastNote := files[0]
		readNote(lastNote)

		if len(files) > 1 {
			yesterdaysNotes := files[1]
			readNote(yesterdaysNotes)
		}

		enterToContinue()
		return
	}

	fmt.Println("\nThere are no available notes")
	enterToContinue()
}

// Parse all notes in the logs directory and display the one by date
func fetchNotesbyDate(reader bufio.Reader) {
	fetchAllNotes()
	fmt.Print(Warn("\nEnter date (yy-mm-dd) or (q) to quit -> "))
	date := getInput(reader)
	if date == "q" {
		return
	}
	filesSet, _ := parseLogsDir()

	var note string
	note = filesSet[date]
	if note != "" {
		readNote(note)
		enterToContinue()
		return
	}

	fmt.Println("There are no available notes")
	enterToContinue()
}

// Parse all notes in the logs directory and return all of their filenames
func fetchAllNotes() {
	_, files := parseLogsDir()

	if len(files) > 0 {
		fmt.Println(Teal("\n--=== Notes ===--\n"))
		for _, fileName := range files {
			fmt.Println(Teal(fileName))
		}
		enterToContinue()
		return
	}

	fmt.Println("There are no available notes")
	enterToContinue()
}

// parse logs and sort them by latest date first,
// returns a map and an array of filenames
func parseLogsDir() (map[string]string, []string) {
	fileInfo, err := ioutil.ReadDir(logPath())
	check(err)

	// store the files in a slice sorted
	var files []string
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	// create a filename map
	filesSet := make(map[string]string)
	for _, file := range fileInfo {
		filesSet[file.Name()] = file.Name()
	}
	return filesSet, files
}

// read a specific note file from the logs folder
func readNote(noteName string) {
	content, err := ioutil.ReadFile(logPath() + "/" + noteName)
	check(err)

	fmt.Println(Succ("\nFile stored at: ") + logPath() + "/" + noteName)
	fmt.Printf(Teal("File contents [%s]:\n \n%s"), noteName, content)
}
