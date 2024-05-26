package main

import (
	"fmt"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
    Save() error
}

type outputtable interface {
    saver
    Display()
}

func outputData(data outputtable) error {
    data.Display()
    return saveData(data)
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

func doSomething(value any) {
    switch v := value.(type) {
    case int: 
        fmt.Println("This is an integer:", v)
    case string:
        fmt.Println("This is a string:", v)
    case float64:
        fmt.Println("This is a base64 encoded string:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    title, content := note.GetNoteData()

    userNote, err := note.New(title, content)

    if err != nil {
        fmt.Println(err)
        return
    }

    userContent := todo.GetTodoData()

    userTodo, err := todo.New(userContent)

    if err != nil {
        fmt.Println(err)
        return
    }

    err = outputData(userNote)

    if err != nil {
        return
    }

    err = outputData(userTodo)

    if err != nil {
        return
    }

    fmt.Println("The note was saved successfully.")
    fmt.Println("The todo was saved successfully.")
}
