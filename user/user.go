package user

import (
	"fmt"
)

func GetUserInput(prompt string) string {
    fmt.Print(prompt)
    var value string
    fmt.Scanln(&value)
    return value
}
