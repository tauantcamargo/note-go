package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/note/user"
)

type Note struct {
    Title     string `json:"title"`
    Content   string `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
    if title == "" || content == "" {
        return nil, errors.New("error: Invalid input, input cannot be empty")
    }

    return &Note{
        Title:     title,
        Content:   content,
        CreatedAt: time.Now(),
    }, nil
}

func (note Note) Display()  {
    fmt.Printf("Your note titled %v has the following content:\n\n%v", note.Title, note.Content)
}

func (note Note) Save() error {
    fileName := strings.ReplaceAll(note.Title, " ", "_")
    fileName = strings.ToLower(fileName) + ".json"

    json, err := json.Marshal(note)
    
    if err != nil {
        return err
    }

    return os.WriteFile(fileName, json, 0644)
}

func GetNoteData() (string, string) {
    title := user.GetUserInput("Enter note title: ")
    
    content := user.GetUserInput("Enter note content: ")

    return title, content
}
