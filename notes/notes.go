package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

func main() {
	title, content := getNoteData()
  todoText := getUserInput("Todo text")

	userNote, err := note.New(title, content)
	if err != nil { panic(err) }
	todo, err := todo.New(todoText)
	if err != nil { panic(err) }

	todo.Display()
	err = todo.Save()
	if err != nil { panic(err) }

	userNote.Display()
	err = userNote.Save()
	if err != nil { panic(err) }
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