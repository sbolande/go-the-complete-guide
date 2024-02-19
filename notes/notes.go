package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type Saver interface {
	Save() error
}
type Outputtable interface {
	Saver
	Display()
}

func main() {
	title, content := getNoteData()
  todoText := getUserInput("Todo text")

	userNote, err := note.New(title, content)
	if err != nil { panic(err) }
	todo, err := todo.New(todoText)
	if err != nil { panic(err) }

	err = outputData(todo)
	if err != nil { return }

	err = outputData(userNote)
	if err != nil { return }
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil { 
		fmt.Println("Saving was unsuccessful.")
		return err
	}
	fmt.Println("Saving was successful!")
	return nil
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