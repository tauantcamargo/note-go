package main

import (
	"fmt"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
    Save() error
}

func saveData(data saver) error {
    err := data.Save()

    if err != nil {
        fmt.Printf("Failed to save data: %v\n", err)
        return err
    }

    fmt.Println("Data saved successfully.")
    return nil
}

func main() {
    title, content := note.GetNoteData()

    userNote, error := note.New(title, content)

    if error != nil {
        fmt.Println(error)
        return
    }

    userContent := todo.GetTodoData()

    userTodo, error := todo.New(userContent)

    if error != nil {
        fmt.Println(error)
        return
    }

    userNote.Display()
    saveData(userNote)

    userTodo.Display()
    saveData(userTodo)
    
    fmt.Println("The note was saved successfully.")
    fmt.Println("The todo was saved successfully.")
}
