package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
    title, content, error := note.GetNoteData()

    if error != nil {
        fmt.Println(error)
        return
    }
}
