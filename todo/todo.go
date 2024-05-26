package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"example.com/note/user"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {
	if (text == "") {
		return nil, errors.New("error: Invalid input, input cannot be empty")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	json, err := json.Marshal(todo)

	if (err != nil) {
		return err
	}

	return os.WriteFile("todo.json", json, 0644)
}

func GetTodoData() string {
	text := user.GetUserInput("Enter todo text: ")

	return text
}