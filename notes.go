package main

import (
	"fmt"

	"example.com/notes/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title")
	content := getUserInput("Note content")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt, ": ")
	var value string
	fmt.Scanln(&value)
	return value
}