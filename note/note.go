package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (n Note) Display() {
	fmt.Printf("Your note title %v has the following content:\n\n%v", n.title, n.content)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}