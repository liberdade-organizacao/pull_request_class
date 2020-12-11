package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)


func toUpper(s string) string {
    return strings.ToUpper(s)
}


func main() {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    upperText := toUpper(text)
    fmt.Println(upperText)
}
