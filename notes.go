package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil { panic(err) }
	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note title")
	content := getUserInput("Note content")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt, ": ")
	
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil { return "" }
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}