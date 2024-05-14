package main

import (
	"fmt"

	"example.com/note/user"
)

func getNoteData() (string, string, error) {
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

func main() {
    title, content, error := getNoteData()

    if error != nil {
        fmt.Println(error)
        return
    }
}
