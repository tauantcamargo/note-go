package note

import (
	"errors"
	"time"

	"example.com/note/user"
)

type Note struct {
    title     string
    content   string
    createdAt time.Time
}

func New(title, content string) (*Note, error) {
    if title == "" {
        return nil, errors.New("error: Title cannot be empty")
    }

    if content == "" {
        return nil, errors.New("error: Content cannot be empty")
    }

    return &Note{
        title:     title,
        content:   content,
        createdAt: time.Now(),
    }, nil
}

func GetNoteData() (string, string, error) {
    title, error := user.GetUserInput("Enter note title: ")
    
    if error != nil {
        return "", "", error
    }
    
    content, error := user.GetUserInput("Enter note content: ")

    if error != nil {
        return "", "", error
    }

    return title, content, nil
}
