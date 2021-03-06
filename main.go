package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

const (
	layoutDate = "06-01-02"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		show_menu()

		fmt.Print("-> ")
		text := getInput(*reader)

		if strings.Compare("1", text) == 0 {
			clearScreen()
			add_note(*reader)
		}
		if strings.Compare("2", text) == 0 {
			clearScreen()
			fetch_note(*reader)
		}
		if strings.Compare("3", text) == 0 {
			clearScreen()
			delete_note(*reader)
		}

		if strings.Compare("q", text) == 0 {
			fmt.Println("Goodbye...")
			os.Exit(1)
		}
		clearScreen()
	}
}

// Display main menu
func show_menu() {
	fmt.Println("")
	fmt.Println(Info("Daily Standup Notes:"))
	fmt.Println("---------------------")
	fmt.Println(Succ("(1)") + " Add notes")
	fmt.Println(Succ("(2)") + " Fetch notes")
	fmt.Println(Succ("(3)") + " Delete notes")
	fmt.Println(Succ("(q)") + " Exit")
	fmt.Println("---------------------")
}

func logPath() string {
	usr, err := user.Current()
	check(err)

	return usr.HomeDir + "/dsu_logs"
}

func enterToContinue() {
	fmt.Print("\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(reader bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	return strings.Replace(text, "\n", "", -1)
}

// Clear the terminal
func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
