package user

import (
	"errors"
	"fmt"
)

func GetUserInput(prompt string) (string, error) {
    fmt.Print(prompt)
    var value string
    fmt.Scanln(&value)
    
    if value == "" {
        return "", errors.New("error: Input cannot be empty")
    }

    return value, nil
}
