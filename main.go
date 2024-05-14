package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
    title, content := note.GetNoteData()

    userNote, error := note.New(title, content)

    if error != nil {
        fmt.Println(error)
        return
    }

    userNote.Display()
    err := userNote.Save()

    if err != nil {
        fmt.Println("Saving the note failed.")
        return
    }

    fmt.Println("The note was saved successfully.")
}
