package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
    fmt.Print(prompt)
    reader := bufio.NewReader(os.Stdin)
    text, error := reader.ReadString('\n')

    if error != nil {
        return ""
    }

    text = strings.TrimSuffix(text, "\n")
    text = strings.TrimSuffix(text, "\r")

    return text
}
