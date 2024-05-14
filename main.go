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
}
