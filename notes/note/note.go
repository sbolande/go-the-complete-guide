package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Display() {
	fmt.Printf("Your note title %v has the following content:\n\n%v", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(n.Title, " ", "_")) + ".json"
	json, err := json.Marshal(n)
	if err != nil { return err }
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}