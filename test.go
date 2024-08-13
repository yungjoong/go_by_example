package main

import (
    "strings"
    "fmt"
)

func CleanUpMessage(oldMsg string) string {
    return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}


func main() {
    message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
    fmt.Println(CleanUpMessage(message))
}
