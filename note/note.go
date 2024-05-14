package note

import (
	"errors"
	"fmt"
	"time"

	"example.com/note/user"
)

type Note struct {
    title     string
    content   string
    createdAt time.Time
}

func New(title, content string) (*Note, error) {
    if title == "" || content == "" {
        return nil, errors.New("error: Invalid input, input cannot be empty")
    }

    return &Note{
        title:     title,
        content:   content,
        createdAt: time.Now(),
    }, nil
}

func (note *Note) Display()  {
    fmt.Printf("Your note titled %v has the following content:\n\n%v", note.title, note.content)
}

func GetNoteData() (string, string) {
    title := user.GetUserInput("Enter note title: ")
    
    content := user.GetUserInput("Enter note content: ")

    return title, content
}
