package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"
)

const (
	layoutDate = "1_Jan_2006"
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

// Display menu text
func show_menu() {
	fmt.Println("")
	fmt.Println("Daily Standup Notes:")
	fmt.Println("---------------------")
	fmt.Println("(1) -> Add a note")
	fmt.Println("(2) -> Fetch a note")
	fmt.Println("(3) -> Delete a note")
	fmt.Println("(q) -> Exit")
	fmt.Println("---------------------")
}

// Prepare to create a note file with a date
func add_note(reader bufio.Reader) {
	for {
		create_home_dir_logs_storage()
		show_add_note_menu()

		fmt.Print("(Add note)-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("1", text) == 0 {
			fmt.Print("(Add note +)-> ")
			note, _ := reader.ReadString('\n')
			// convert CRLF to LF
			note = strings.Replace(note, "\n", "", -1)
			add(note, time.Now())
		}

		if strings.Compare("q", text) == 0 {
			break
		}
	}
}

// Display add note menu text
func show_add_note_menu() {
	fmt.Println("")
	fmt.Println("Add a note:")
	fmt.Println("---------------------")
	fmt.Println("(1) -> Add a note for today")
	fmt.Println("(2) -> Add a note for another day")
	fmt.Println("(q) -> Back")
	fmt.Println("---------------------")
}

// Add a note, creates a filename with the given date as a name
// and it is stored under the logs folder
func add(note string, t time.Time) {
	usr, err := user.Current()
	check(err)
	filepath := fmt.Sprintf("%s/dsu_logs/%s", usr.HomeDir, t.Format(layoutDate))

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()

	if _, err = f.WriteString("• " + note + "\n"); err != nil {
		panic(err)
	}
	fmt.Printf("Added: %s\n", note)
	f.Sync()
}

// Create a folder that holds all the logs in the home directory
func create_home_dir_logs_storage() {
	usr, err := user.Current()
	check(err)

	destinationDir := usr.HomeDir + "/dsu_logs"
	_ = os.Mkdir(destinationDir, 0700)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
